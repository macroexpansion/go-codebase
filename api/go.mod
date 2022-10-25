module api

go 1.19

require api/lib v0.0.0
require pgsql v0.0.0
require github.com/gorilla/mux v1.8.0 // indirect

replace api/lib v0.0.0 => ./src/lib
replace pgsql v0.0.0 => ../database/src/lib/pgsql
