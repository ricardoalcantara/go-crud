package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           DbUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()));"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Name         *string `gorm:"type:varchar(255);not null"`
	Email        *string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Username     *string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password     *string `gorm:"type:varchar(255);not null"`
	LastLogin    *time.Time
	FailAttempts string `gorm:"type:tinyint;not null;default:(0)"`
	LastAttempt  *time.Time

	UserPasswordRecoverys []UserPasswordRecovery
}
