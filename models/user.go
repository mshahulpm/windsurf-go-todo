package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Role type for user roles
type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

// User represents a user in the system
type User struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Username  string     `json:"username" gorm:"unique;not null" example:"johndoe"`
	Password  string     `json:"-" example:"secretpassword"`
	Email     string     `json:"email" gorm:"unique;not null" example:"john@example.com"`
	Role      Role       `json:"role" gorm:"type:varchar(20);default:'user'" example:"user"`
	Todos     []Todo     `json:"todos,omitempty"`
	LastLogin time.Time  `json:"last_login,omitempty"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
