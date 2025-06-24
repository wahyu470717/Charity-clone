package services

import (
	"errors"
	"fmt"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/models"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepositoryInterface
	roleRepo repository.RoleRepositoryInterface
	jwtUtil  utils.JWTUtilInterface
}

func NewAuthService(
	userRepo repository.UserRepositoryInterface,
	roleRepo repository.RoleRepositoryInterface,
	jwtUtil utils.JWTUtilInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *AuthService) Login(email, password string) (*response.SignInResponse, error) {
	// Get user by username
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "email not found" {
			return nil, errors.New("email not found")
		}
		return nil, fmt.Errorf("failed to get email: %v", err)
	}

	// Check password
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	// Generate JWT token
	// Dapatkan nama role
	role, err := s.roleRepo.GetRoleByID(user.RoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %v", err)
	}

	token, err := s.jwtUtil.GenerateJWT(user.Username, int64(user.UserID), string(user.RoleID))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	response := &response.SignInResponse{
		UserID:   int64(user.UserID),
		UserName: user.Username,
		Role:     role.RoleName,
		Token:    token,
	}

	return response, nil
}

func (s *AuthService) Register(req request.RegisterRequest) (*response.RegisterResponse, error) {
	// Check if username already exists
	usernameExists, err := s.userRepo.CheckUsernameExists(req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username: %v", err)
	}
	if usernameExists {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	emailExists, err := s.userRepo.CheckEmailExists(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %v", err)
	}
	if emailExists {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create user model
	user := &models.User{
		Username:    req.Username,
		Fullname:    req.Fullname,
		Email:       req.Email,
		Password:    string(hashedPassword),
		RoleID:      req.Role,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		IsActive:    true,
	}

	// Create user in database
	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	response := &response.RegisterResponse{
		UserID:   int64(createdUser.UserID),
		UserName: createdUser.Username,
		Email:    createdUser.Email,
		Fullname: createdUser.Fullname,
		RoleID:   createdUser.RoleID,
	}

	return response, nil
}

func (s *AuthService) ForgetPassword(email string) (*response.Meta, error) {
	// Check if user exists
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, errors.New("email not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// Generate reset token
	jwtUtil := s.jwtUtil.(*utils.JWTUtil) // Type assertion
	resetToken, err := jwtUtil.GenerateResetToken(int64(user.UserID), user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate reset token: %v", err)
	}

	// TODO: Send email with reset token
	// For now, we'll just return success
	// In production, you would send an email with the reset link
	fmt.Printf("Reset token for %s: %s\n", email, resetToken)

	return &response.Meta{
		Code:    200,
		Message: "Password reset email sent successfully",
		Status:  "success",
	}, nil
}

func (s *AuthService) ChangePassword(userID int64, newPassword string) error {
	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Update password in database
	err = s.userRepo.UpdateUserPassword(userID, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}
