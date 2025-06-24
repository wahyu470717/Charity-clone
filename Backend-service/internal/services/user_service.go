package services

import (
	"context"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID int64) (*response.UserProfileResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &response.UserProfileResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		Fullname:       user.Fullname,
		Email:          user.Email,
		RoleID:         user.RoleID,
		ProfilePicture: user.ProfilePicture,
		PhoneNumber:    user.PhoneNumber,
		Address:        user.Address,
	}, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID int64, req request.UpdateProfileRequest) (*response.UserProfileResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.ProfilePicture != "" {
		user.ProfilePicture = req.ProfilePicture
	}

	// Update in database
	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &response.UserProfileResponse{
		UserID:         user.UserID,
		Username:       user.Username,
		Fullname:       user.Fullname,
		Email:          user.Email,
		RoleID:         user.RoleID,
		ProfilePicture: user.ProfilePicture,
		PhoneNumber:    user.PhoneNumber,
		Address:        user.Address,
	}, nil
}
