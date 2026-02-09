package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidPassword   = errors.New("invalid password")
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// Create creates a new user
func (m *UserModel) Create(user *User) error {
	return m.db.Create(user).Error
}

// FindByUsername finds a user by username
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

// FindByEmail finds a user by email
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

// FindByID finds a user by ID
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

// Update updates user information
func (m *UserModel) Update(user *User) error {
	return m.db.Save(user).Error
}

// UpdatePassword updates user password
func (m *UserModel) UpdatePassword(userID int64, newPassword string) error {
	return m.db.Model(&User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// CheckUsernameExists checks if username already exists
func (m *UserModel) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := m.db.Model(&User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// CheckEmailExists checks if email already exists
func (m *UserModel) CheckEmailExists(email string) (bool, error) {
	var count int64
	err := m.db.Model(&User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
