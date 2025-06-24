package handlers

import (
	// "context"
	"net/http"
	"share-the-meal/internal/config"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/services"
	"share-the-meal/internal/utils"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AuthHandler struct {
	DB      *pgxpool.Pool
	Logger  *zap.Logger
	service *services.AuthService
	jwtUtil *utils.JWTUtil
}

func NewAuthHandler(db *pgxpool.Pool, logger *zap.Logger) *AuthHandler {
	userRepo := repository.NewUserRepository(db, "public")
	roleRepo := repository.NewRoleRepository(db, "public")
	cfg, _ := config.GetConfig()
	jwtUtil := utils.NewJWTUtil(cfg.JWTSecret)
	authService := services.NewAuthService(userRepo, roleRepo, jwtUtil)

	return &AuthHandler{
		DB:      db,
		Logger:  logger,
		service: authService,
		jwtUtil: jwtUtil,
	}
}

func (h *AuthHandler) SignInUser(c *gin.Context) {
	h.Logger.Info("Attempting to sign in...")

	var req request.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.Meta{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Status:  "error",
		})
		return
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		var statusCode int
		var errorMessage string

		if err.Error() == "user not found" {
			statusCode = http.StatusNotFound
			errorMessage = "User not found"
		} else if err.Error() == "password salah" {
			statusCode = http.StatusUnauthorized
			errorMessage = "Incorrect password"
		} else {
			statusCode = http.StatusInternalServerError
			errorMessage = "Internal server error"
		}

		h.Logger.Error("Login failed", zap.String("username", req.Email), zap.Error(err))
		c.JSON(statusCode, response.Meta{
			Code:    statusCode,
			Message: errorMessage,
			Status:  http.StatusText(statusCode),
		})
		return
	}

	apiResponse := response.APIResponse{
		Data: user,
		Meta: response.Meta{
			Code:    http.StatusOK,
			Message: "Sign in successful",
			Status:  http.StatusText(http.StatusOK),
		},
	}

	c.JSON(http.StatusOK, apiResponse)
}

// POST /api/v1/auth-management/register
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	h.Logger.Info("Attempting to register user...")

	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.Meta{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Status:  "error",
		})
		return
	}

	user, err := h.service.Register(req)
	if err != nil {
		var statusCode int
		var errorMessage string

		if err.Error() == "username already exists" {
			statusCode = http.StatusConflict
			errorMessage = "Username already exists"
		} else if err.Error() == "email already exists" {
			statusCode = http.StatusConflict
			errorMessage = "Email already exists"
		} else {
			statusCode = http.StatusInternalServerError
			errorMessage = "Internal server error"
		}

		h.Logger.Error("Registration failed", zap.String("username", req.Username), zap.Error(err))
		c.JSON(statusCode, response.Meta{
			Code:    statusCode,
			Message: errorMessage,
			Status:  http.StatusText(statusCode),
		})
		return
	}

	apiResponse := response.APIResponse{
		Data: user,
		Meta: response.Meta{
			Code:    http.StatusOK,
			Message: "Sign in successful",
			Status:  http.StatusText(http.StatusOK),
		},
	}

	c.JSON(http.StatusCreated, apiResponse)
}

// POST /api/v1/auth-management/forgot-password
func (h *AuthHandler) ForgetPassword(c *gin.Context) {
	var req request.ForgetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Forget password validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.Meta{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Status:  "error",
		})
		return
	}

	result, err := h.service.ForgetPassword(req.Email)
	if err != nil {
		var statusCode int
		var errorMessage string

		if err.Error() == "email not found" {
			statusCode = http.StatusNotFound
			errorMessage = "Email not found"
		} else {
			statusCode = http.StatusInternalServerError
			errorMessage = "Internal server error"
		}

		h.Logger.Error("Failed to process forget password request", zap.Error(err))
		c.JSON(statusCode, response.Meta{
			Code:    statusCode,
			Message: errorMessage,
			Status:  http.StatusText(statusCode),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// PUT /api/v1/auth-management/change-password
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		h.Logger.Error("Token is required")
		c.JSON(http.StatusBadRequest, response.Meta{
			Code:    http.StatusBadRequest,
			Message: "Token is required",
			Status:  "error",
		})
		return
	}

	var req request.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.Meta{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Status:  "error",
		})
		return
	}

	claims, err := h.jwtUtil.VerifyToken(token)
	if err != nil {
		h.Logger.Error("Invalid token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, response.Meta{
			Code:    http.StatusUnauthorized,
			Message: "Invalid or expired token",
			Status:  "error",
		})
		return
	}

	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		h.Logger.Error("Invalid user ID in token")
		c.JSON(http.StatusBadRequest, response.Meta{
			Code:    http.StatusBadRequest,
			Message: "Invalid user ID in token",
			Status:  "error",
		})
		return
	}

	userID := int64(userIDFloat)

	err = h.service.ChangePassword(userID, req.Password)
	if err != nil {
		h.Logger.Error("Failed to change password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.Meta{
			Code:    http.StatusInternalServerError,
			Message: "Failed to change password",
			Status:  "error",
		})
		return
	}

	c.JSON(http.StatusOK, response.Meta{
		Code:    http.StatusOK,
		Message: "Password changed successfully",
		Status:  "success",
	})
}
