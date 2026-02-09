package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrPasswordNotFound = errors.New("password record not found")
)

type PasswordModel struct {
	db *gorm.DB
}

func NewPasswordModel(db *gorm.DB) *PasswordModel {
	return &PasswordModel{db: db}
}

// Create creates a new password record
func (m *PasswordModel) Create(password *Password) error {
	return m.db.Create(password).Error
}

// FindByID finds a password record by ID
func (m *PasswordModel) FindByID(id int64) (*Password, error) {
	var password Password
	err := m.db.Where("id = ?", id).First(&password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPasswordNotFound
		}
		return nil, err
	}
	return &password, nil
}

// FindByUserID finds all password records by user ID
func (m *PasswordModel) FindByUserID(userID int64) ([]*Password, error) {
	var passwords []*Password
	err := m.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&passwords).Error
	if err != nil {
		return nil, err
	}
	return passwords, nil
}

// Update updates a password record
func (m *PasswordModel) Update(password *Password) error {
	return m.db.Save(password).Error
}

// Delete deletes a password record
func (m *PasswordModel) Delete(id int64, userID int64) error {
	result := m.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Password{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPasswordNotFound
	}
	return nil
}

// CheckOwnership checks if a password record belongs to the specified user
func (m *PasswordModel) CheckOwnership(id int64, userID int64) (bool, error) {
	var count int64
	err := m.db.Model(&Password{}).Where("id = ? AND user_id = ?", id, userID).Count(&count).Error
	return count > 0, err
}
