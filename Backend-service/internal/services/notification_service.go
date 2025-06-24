package services

import (
	"share-the-meal/internal/models"
	"share-the-meal/internal/repository"
)

type NotificationService struct {
	repo *repository.NotificationsRepository
}

func NewNotificationService(repo *repository.NotificationsRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) CreateNotification(userID int64, title, message string) error {
	notification := &models.Notifications{
		UserID:  userID,
		Title:   title,
		Message: message,
		IsRead:  false,
	}
	
	return s.repo.CreateNotification(notification)
}

func (s *NotificationService) GetUserNotifications(userID int64) ([]models.Notifications, error) {
	return s.repo.GetUserNotifications(userID)
}