package model

import (
	"time"
)

// Password represents the password model
type Password struct {
	ID          int64     `gorm:"primarykey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	UserID      int64     `gorm:"index;not null" json:"user_id"`
	Password    string    `gorm:"type:text;not null" json:"-"` // Password is not returned to client, needs separate handling
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name
func (Password) TableName() string {
	return "passwords"
}
