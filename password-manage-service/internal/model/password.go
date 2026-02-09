package model

import (
	"time"
)

// Password 密码模型
type Password struct {
	ID          int64     `gorm:"primarykey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	UserID      int64     `gorm:"index;not null" json:"user_id"`
	Password    string    `gorm:"type:text;not null" json:"-"` // 密码不返回给客户端，需要单独处理
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Password) TableName() string {
	return "passwords"
}
