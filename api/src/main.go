package main

import (
	"fmt"

	"api/lib"
	"api/lib/hello_lib"
	"api/src/internal"
	"api/src/internal/controllers"

	"pgsql"
)

func main() {
	fmt.Println("test")

	internal.Hello()
	hello_lib.Hello()

	controllers.Login()
	lib.External()

	pgsql.Hello()
}
