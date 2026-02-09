package model

import (
	"time"
)

// Website represents the website model
type Website struct {
	ID          int64     `gorm:"primarykey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Icon        string    `gorm:"size:500" json:"icon"` // Icon URL or base64
	Description string    `gorm:"type:text" json:"description"`
	URL         string    `gorm:"size:500;not null" json:"url"`
	UserID      int64     `gorm:"index;not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name
func (Website) TableName() string {
	return "websites"
}
