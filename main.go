package main

import (
	"fmt"
	_ "github.com/dabfleming/go-martini-example/controllers"
	"github.com/dabfleming/go-martini-example/server"
)

func main() {
	fmt.Println("main.go main()")
	server.Server.Run()
}
