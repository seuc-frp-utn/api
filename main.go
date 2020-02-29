package main

import "github.com/seuc-frp-utn/api/database"

func init() {

}

func main() {
	router := initializeRouter()
	database.SetupDatabase()
	router.Run(":8080")
}
