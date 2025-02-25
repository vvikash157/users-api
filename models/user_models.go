package models

import "time"

type User struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      string    `json:"user_id" gorm:"unique;not null"`
	AccessToken string    `json:"access_token"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	Age         int       `json:"age"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically set when creating
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Automatically update on modification
}
