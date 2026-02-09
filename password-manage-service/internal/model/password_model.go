package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrPasswordNotFound = errors.New("密码记录不存在")
)

type PasswordModel struct {
	db *gorm.DB
}

func NewPasswordModel(db *gorm.DB) *PasswordModel {
	return &PasswordModel{db: db}
}

// Create 创建密码记录
func (m *PasswordModel) Create(password *Password) error {
	return m.db.Create(password).Error
}

// FindByID 根据ID查找密码记录
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

// FindByUserID 根据用户ID查找所有密码记录
func (m *PasswordModel) FindByUserID(userID int64) ([]*Password, error) {
	var passwords []*Password
	err := m.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&passwords).Error
	if err != nil {
		return nil, err
	}
	return passwords, nil
}

// Update 更新密码记录
func (m *PasswordModel) Update(password *Password) error {
	return m.db.Save(password).Error
}

// Delete 删除密码记录
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

// CheckOwnership 检查密码记录是否属于指定用户
func (m *PasswordModel) CheckOwnership(id int64, userID int64) (bool, error) {
	var count int64
	err := m.db.Model(&Password{}).Where("id = ? AND user_id = ?", id, userID).Count(&count).Error
	return count > 0, err
}
