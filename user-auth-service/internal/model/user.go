package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        int64     `gorm:"primarykey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"` // 密码不返回给客户端
	Email     string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Mobile    string    `gorm:"size:20" json:"mobile"`
	Status    int       `gorm:"default:1" json:"status"` // 1:正常 0:禁用
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
