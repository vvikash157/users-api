package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      string `json:"user_id" gorm:"unique;not null"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	Age         int    `json:"age"`
}
