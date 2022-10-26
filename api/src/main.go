package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"api/src/internal/router"
	"pgsql"
	"redis"
)

func main() {
	// .env
	env := os.Getenv("APP_ENV")
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// pgsql
	pg, err := pgsql.Connect(os.Getenv("PGSQL_HOST"), os.Getenv("PGSQL_USER"), os.Getenv("PGSQL_PASSWORD"), os.Getenv("PGSQL_DBNAME"))
	if err != nil {
		log.Fatal("Error connecting to pgSQL")
	}
	pgsql.Migrate(pg)
	log.Println("pgSQL connected")

	// redis
	redis, err := redis.Connect(os.Getenv("REDIS_ADDR"))
	_ = redis
	if err != nil {
		log.Fatal("Error connecting Redis")
	}
	log.Println("Redis connected")

	// mux
	http.Handle("/", router.Router(pg))

	log.Println("Listen on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
