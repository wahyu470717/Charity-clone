package repository

import (
	"context"
	"fmt"
	"share-the-meal/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NotificationsRepositoryInterface interface {
	CreateNotification(notification *models.Notifications) error
	GetUserNotifications(userID int64) ([]models.Notifications, error)
}

type NotificationsRepository struct {
	db     *pgxpool.Pool
	schema string
}

func NewNotificationsRepository(db *pgxpool.Pool, schema string) *NotificationsRepository {
	return &NotificationsRepository{
		db:     db,
		schema: schema,
	}
}

func (r *NotificationsRepository) CreateNotification(notification *models.Notifications) error {
	query := `
		INSERT INTO notifications (user_id, title, message, is_read, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(context.Background(), query,
		notification.UserID,
		notification.Title,
		notification.Message,
		notification.IsRead,
		time.Now(),
	).Scan(&notification.ID)

	if err != nil {
		return fmt.Errorf("failed to create notification: %v", err)
	}

	return nil
}

func (r *NotificationsRepository) GetUserNotifications(userID int64) ([]models.Notifications, error) {
	query := `
		SELECT id, user_id, title, message, is_read, created_at
		FROM notifications
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %v", err)
	}
	defer rows.Close()

	var notifications []models.Notifications
	for rows.Next() {
		var n models.Notifications
		err := rows.Scan(
			&n.ID,
			&n.UserID,
			&n.Title,
			&n.Message,
			&n.IsRead,
			&n.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}

	return notifications, nil
}
