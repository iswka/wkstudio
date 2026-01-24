package model

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound      = errors.New("用户不存在")
	ErrUserAlreadyExists = errors.New("用户已存在")
	ErrInvalidPassword   = errors.New("密码错误")
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// Create 创建用户
func (m *UserModel) Create(user *User) error {
	return m.db.Create(user).Error
}

// FindByUsername 根据用户名查找
func (m *UserModel) FindByUsername(username string) (*User, error) {
	var user User
	err := m.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找
func (m *UserModel) FindByEmail(email string) (*User, error) {
	var user User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// FindByID 根据ID查找
func (m *UserModel) FindByID(id int64) (*User, error) {
	var user User
	err := m.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (m *UserModel) Update(user *User) error {
	return m.db.Save(user).Error
}

// UpdatePassword 更新密码
func (m *UserModel) UpdatePassword(userID int64, newPassword string) error {
	return m.db.Model(&User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// CheckUsernameExists 检查用户名是否存在
func (m *UserModel) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := m.db.Model(&User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// CheckEmailExists 检查邮箱是否存在
func (m *UserModel) CheckEmailExists(email string) (bool, error) {
	var count int64
	err := m.db.Model(&User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
