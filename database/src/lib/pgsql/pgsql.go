package pgsql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pgsql/models"
)

func Hello() {
	println("Hello pgSQL")
}

func Connect(host string, user string, password string, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Account{})
}
