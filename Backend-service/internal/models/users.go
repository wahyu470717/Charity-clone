package models

import "time"

type User struct {
	UserID         int64     `json:"id" db:"user_id"`
	Username       string    `json:"username" db:"username"`
	Fullname       string    `json:"fullname" db:"fullname"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"password"`
	RoleID         int64     `json:"role_id" db:"role_id"`
	ProfilePicture string    `json:"profile_picture,omitempty" db:"profile_picture"`
	PhoneNumber    string    `json:"phone_number,omitempty" db:"phone_number"`
	Address        string    `json:"address,omitempty" db:"address"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	ModifiedBy     string    `json:"modified_by" db:"modified_by"`
	ModifiedAt     time.Time `json:"modified_at" db:"modified_at"`
}
