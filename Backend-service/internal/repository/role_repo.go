package repository

import (
	"context"
	"fmt"
	"share-the-meal/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepositoryInterface interface {
	GetRoleByID(roleID int64) (*models.Role, error)
}

type RoleRepository struct {
	db     *pgxpool.Pool
	schema string
}

func NewRoleRepository(db *pgxpool.Pool, schema string) *RoleRepository {
	return &RoleRepository{
		db:     db,
		schema: schema,
	}
}

func (r *RoleRepository) GetRoleByID(roleID int64) (*models.Role, error) {
	query := `
		SELECT 
			role_id, 
			role_name, 
			role_description,
			created_at,
			modified_at
		FROM roles 
		WHERE role_id = $1
	`

	var role models.Role

	err := r.db.QueryRow(context.Background(), query, roleID).Scan(
		&role.RoleID,
		&role.RoleName,
		&role.RoleDescription,
		&role.CreatedAt,
		&role.ModifiedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("role not found")
		}
		return nil, fmt.Errorf("failed to get role: %v", err)
	}

	return &role, nil
}

func (r *RoleRepository) GetRoleByName(roleName string) (*models.Role, error) {

	query := `
		SELECT 
			role_id, 
			role_name, 
			role_description
		FROM roles
		WHERE name = $1 AND
		 is_active = true
		`

	var role models.Role
	err := r.db.QueryRow(context.Background(), query, roleName).Scan(
		&role.RoleID,
		&roleName,
		&role.RoleDescription,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("role not found")
		}
		return nil, fmt.Errorf("failed to get role: %v", err)
	}

	return &role, nil
}

func (r *RoleRepository) CreateRole(role *models.Role) error {
	query := `
		INSERT INTO roles (role_name, role_description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING role_id
	`

	now := time.Now()
	err := r.db.QueryRow(context.Background(), query,
		role.RoleName,
		role.RoleDescription,
		now,
		now,
	).Scan(&role.RoleID)
	if err != nil {
		return fmt.Errorf("failed to create role: %v", err)
	}

	return nil
}
