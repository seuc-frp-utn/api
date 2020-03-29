package main

import "github.com/seuc-frp-utn/api/pkg/database"

func main() {
	database.SetupDatabase()
	defer database.Close()
	router := initializeRouter()
	router.Run(":3000")
}
