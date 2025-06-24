package handlers

import (
    "net/http"
    "share-the-meal/internal/dto/response"
    "share-the-meal/internal/repository"
    "share-the-meal/internal/services"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
    "go.uber.org/zap"
)

type UserHandler struct {
    userService *services.UserService
    logger      *zap.Logger
}

func NewUserHandler(db *pgxpool.Pool, logger *zap.Logger) *UserHandler {
    userRepo := repository.NewUserRepository(db, "public")
    return &UserHandler{
        userService: services.NewUserService(userRepo),
        logger:      logger,
    }
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
    userID := c.MustGet("userID").(int64)
    user, err := h.userService.GetUserProfile(c.Request.Context(), userID)
    if err != nil {
        h.logger.Error("Failed to get user profile", zap.Error(err))
        c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get profile"))
        return
    }
    c.JSON(http.StatusOK, response.SuccessResponse(user))
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
    // Implementasi update profile
}