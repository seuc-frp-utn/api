package main

import "github.com/seuc-frp-utn/api/database"

func main() {
	router := initializeRouter()
	database.SetupDatabase()
	router.Run(":3000")
}
