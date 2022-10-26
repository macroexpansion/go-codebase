package main

import (
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"

	"pgsql"
	pgmodels "pgsql/models"
	"api/src/internal/router"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	pgsql.Hello()
	/* db := pgsql.Connect(os.Getenv("PGSQL_HOST"), os.Getenv("PGSQL_USER"), os.Getenv("PGSQL_PASSWORD"), os.Getenv("PGSQL_DBNAME"))
	_ = db */

	account := pgmodels.Account{
		Username: "quang",
		Password: "qwe123",
	}
	println(account.Username)

	http.Handle("/", router.Router())

	log.Println("Listen on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
