package main

import (
	"fmt"

	"lib"
	"lib/hello_lib"
)

func main() {
	fmt.Println("test")

	lib.External()
	hello_lib.Hello()
}
