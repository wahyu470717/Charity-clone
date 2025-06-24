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

type NotificationHandler struct {
    notificationService *services.NotificationService
    logger              *zap.Logger
}

func NewNotificationHandler(db *pgxpool.Pool, logger *zap.Logger) *NotificationHandler {
    notificationRepo := repository.NewNotificationsRepository(db, "public")
    return &NotificationHandler{
        notificationService: services.NewNotificationService(notificationRepo),
        logger:              logger,
    }
}

func (h *NotificationHandler) GetUserNotifications(c *gin.Context) {
    userID := c.MustGet("userID").(int64)
    notifications, err := h.notificationService.GetUserNotifications(userID)
    if err != nil {
        h.logger.Error("Failed to get notifications", zap.Error(err))
        c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get notifications"))
        return
    }
    c.JSON(http.StatusOK, response.SuccessResponse(notifications))
}