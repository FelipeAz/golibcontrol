package model

import (
	"time"
)

// Account contains all Book's table properties.
type Account struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	Password  string    `json:"password" binding:"required"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}