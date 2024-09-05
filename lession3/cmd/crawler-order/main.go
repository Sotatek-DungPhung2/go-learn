package main

import (
	"fmt"

	authMid "lession3/middlewares"
)

func main() {

	authMid.AuthMiddleWare()
	fmt.Println("Hello World from crawler-order!")
}
