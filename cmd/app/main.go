package main

import "github.com/seuc-frp-utn/api/pkg/database"

func main() {
	database.SetupDatabase()
	router := initializeRouter()
	router.Run(":3000")
}
