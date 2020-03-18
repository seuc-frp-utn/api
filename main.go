package main

import "github.com/seuc-frp-utn/api/database"

func main() {
	database.SetupDatabase()
	router := initializeRouter()
	router.Run(":3000")
}
