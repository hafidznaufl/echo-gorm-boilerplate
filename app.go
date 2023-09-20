package main

import (
	"app/config"
	"app/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
