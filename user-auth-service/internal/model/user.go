package model

import (
	"time"
)

// User represents the user model
type User struct {
	ID        int64     `gorm:"primarykey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"` // Password is not returned to client
	Email     string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Mobile    string    `gorm:"size:20" json:"mobile"`
	Status    int       `gorm:"default:1" json:"status"` // 1:active 0:disabled
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name
func (User) TableName() string {
	return "users"
}
