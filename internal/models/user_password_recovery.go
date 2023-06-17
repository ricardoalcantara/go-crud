package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPasswordRecovery struct {
	ID        DbUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()));"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	UserID    DbUUID `gorm:"not null;index:idx_user_id_hash_code"`
	HashCode  string `gorm:"type:varchar(255);not null;index:idx_user_id_hash_code"`
	Active    bool   `gorm:"not null"`
	Attempt   string `gorm:"type:tinyint;not null;default:(0)"`

	User User
}
