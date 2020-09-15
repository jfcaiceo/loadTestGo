package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID                   uint             `json:"id"`
	Email                string           `json:"email" valid:"required,email"`
	CreatedAt            time.Time        `json:"created_at"`
	UpdatedAt            time.Time        `json:"updated_at"`
	EncryptedPassword    string           `json:"encrypted_password"`
}

func (user *User) All(db *gorm.DB, params map[string]interface {}) ([]*User, error) {	
	users := []*User{}

	if result := db.Where(params).Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (user *User) Get(db *gorm.DB) error {
	if result := db.First(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

