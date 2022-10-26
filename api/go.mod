module api

go 1.19

require api/lib v0.0.0

require pgsql v0.0.0

require redis v0.0.0

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v9 v9.0.0-rc.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/postgres v1.4.5 // indirect
	gorm.io/gorm v1.24.1-0.20221019064659-5dd2bb482755 // indirect
)

replace api/lib v0.0.0 => ./src/lib

replace pgsql v0.0.0 => ../database/src/lib/pgsql

replace redis v0.0.0 => ../database/src/lib/redis
