package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:50;not null"`
	Email     string `gorm:"size:50;not null;unique"`
	Password  string `gorm:"size:50;not null"`
	Address   string `gorm:"size:100;not null"`
	CreatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return nil
}
