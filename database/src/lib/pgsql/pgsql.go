package pgsql

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func Hello() {
	println("Hello pgSQL")
}

func connect(host string, user string, password string, db string) {
	dsn := "host=localhost user=postgres password=qwe123 dbname=dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
