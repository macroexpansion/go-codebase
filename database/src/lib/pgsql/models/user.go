package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string
	Password string
}
