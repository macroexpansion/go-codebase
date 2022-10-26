package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}
