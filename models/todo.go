package models

import (

	"time"
)

// Todo represents a todo item
// @Description Todo represents a todo item in the system
type Todo struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Title       string    `json:"title" binding:"required" example:"Buy groceries"`
	Description string    `json:"description" example:"Buy milk, eggs, and bread"`
	Completed   bool      `json:"completed" gorm:"default:false" example:"false"`
	UserID      uint      `json:"user_id" example:"1"`
	User        User      `json:"user,omitempty" gorm:"foreignkey:UserID"`
}
