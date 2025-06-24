package models

import "time"

type Role struct {
	RoleID          int64     `json:"role_id" db:"role_id"`
	RoleName        string    `json:"role_name" db:"role_name"`
	RoleDescription string    `json:"role_description" db:"role_description"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	ModifiedAt      time.Time `json:"modified_at" db:"modified_at"`
}
