package pgsql

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
	"log"
)

func Hello() {
	println("Hello pgSQL")
}

func Connect(host string, user string, password string, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to pgSQL")
	}
	return db
}
