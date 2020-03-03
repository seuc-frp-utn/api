package main

import "github.com/seuc-frp-utn/api/database"

func main() {
	database.SetupDatabaseTests()
	router := initializeRouter()
	router.Run(":3000")
}
