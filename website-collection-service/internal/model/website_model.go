package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrWebsiteNotFound = errors.New("website not found")
)

type WebsiteModel struct {
	db *gorm.DB
}

func NewWebsiteModel(db *gorm.DB) *WebsiteModel {
	return &WebsiteModel{db: db}
}

// Create creates a new website record
func (m *WebsiteModel) Create(website *Website) error {
	return m.db.Create(website).Error
}

// FindByID finds a website record by ID
func (m *WebsiteModel) FindByID(id int64) (*Website, error) {
	var website Website
	err := m.db.Where("id = ?", id).First(&website).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWebsiteNotFound
		}
		return nil, err
	}
	return &website, nil
}

// FindByUserID finds all website records by user ID
func (m *WebsiteModel) FindByUserID(userID int64) ([]*Website, error) {
	var websites []*Website
	err := m.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&websites).Error
	if err != nil {
		return nil, err
	}
	return websites, nil
}

// Update updates a website record
func (m *WebsiteModel) Update(website *Website) error {
	return m.db.Save(website).Error
}

// Delete deletes a website record
func (m *WebsiteModel) Delete(id int64, userID int64) error {
	result := m.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Website{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrWebsiteNotFound
	}
	return nil
}

// CheckOwnership checks if a website record belongs to the specified user
func (m *WebsiteModel) CheckOwnership(id int64, userID int64) (bool, error) {
	var count int64
	err := m.db.Model(&Website{}).Where("id = ? AND user_id = ?", id, userID).Count(&count).Error
	return count > 0, err
}
