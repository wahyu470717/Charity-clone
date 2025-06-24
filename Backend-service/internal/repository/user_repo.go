package repository

import (
	"context"
	"fmt"
	"share-the-meal/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	GetUserByName(userName string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUserPassword(userID int64, hashedPassword string) error
	CheckUsernameExists(username string) (bool, error)
	CheckEmailExists(email string) (bool, error)
	GetUserByID(ctx context.Context, userID int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type UserRepository struct {
	db     *pgxpool.Pool
	schema string
}

func NewUserRepository(db *pgxpool.Pool, schema string) *UserRepository {
	return &UserRepository{
		db:     db,
		schema: schema,
	}
}

func (r *UserRepository) GetUserByName(userName string) (*models.User, error) {
	query := `
		SELECT 
			user_id, 
			username, 
			role_id, 
			password 
		FROM users 
		WHERE username = $1 AND is_active = true
	`

	var user models.User

	err := r.db.QueryRow(context.Background(), query, userName).Scan(
		&user.UserID,
		&user.Username,
		&user.RoleID,
		&user.Password,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT 
			user_id, 
			username, 
			fullname,
			email,
			password,
			role_id,
			profile_picture,
			phone_number,
			address,
			is_active,
			created_at,
			modified_at
		FROM users 
		WHERE email = $1 AND is_active = true
	`

	var user models.User
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.UserID,
		&user.Username,
		&user.Fullname,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.ProfilePicture,
		&user.PhoneNumber,
		&user.Address,
		&user.IsActive,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `
        INSERT INTO users (
            username, fullname, email, password, role_id, 
            phone_number, address, created_by, created_at, 
            modified_by, modified_at, is_active
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
        ) RETURNING user_id
    ` // Hanya return user_id

	now := time.Now()
	user.CreatedAt = now
	user.ModifiedAt = now
	user.CreatedBy = user.Username
	user.ModifiedBy = user.Username

	// Eksekusi query
	err := r.db.QueryRow(context.Background(), query,
		user.Username,
		user.Fullname,
		user.Email,
		user.Password,
		user.RoleID,
		user.PhoneNumber,
		user.Address,
		user.CreatedBy,
		user.CreatedAt,
		user.ModifiedBy,
		user.ModifiedAt,
		user.IsActive,
	).Scan(&user.UserID)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

func (r *UserRepository) UpdateUserPassword(userID int64, hashedPassword string) error {
	query := `
		UPDATE users 
		SET password = $1, modified_at = $2 
		WHERE user_id = $3 AND is_active = true
	`

	_, err := r.db.Exec(context.Background(), query, hashedPassword, time.Now(), userID)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 AND is_active = true)`

	var exists bool
	err := r.db.QueryRow(context.Background(), query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check username: %v", err)
	}

	return exists, nil
}

func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND is_active = true)`

	var exists bool
	err := r.db.QueryRow(context.Background(), query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email: %v", err)
	}

	return exists, nil
}

// Tambahkan method baru
func (r *UserRepository) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `
		SELECT 
			user_id, 
			username, 
			fullname,
			email,
			password,
			role_id,
			profile_picture,
			phone_number,
			address,
			is_active,
			created_at,
			modified_at
		FROM users 
		WHERE user_id = $1 AND is_active = true
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&user.UserID,
		&user.Username,
		&user.Fullname,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.ProfilePicture,
		&user.PhoneNumber,
		&user.Address,
		&user.IsActive,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET fullname = $1, 
			phone_number = $2, 
			address = $3, 
			profile_picture = $4,
			modified_at = $5
		WHERE user_id = $6
	`

	_, err := r.db.Exec(ctx, query,
		user.Fullname,
		user.PhoneNumber,
		user.Address,
		user.ProfilePicture,
		time.Now(),
		user.UserID,
	)
	return err
}
