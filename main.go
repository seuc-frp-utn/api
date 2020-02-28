package main

func init() {

}

func main() {
	router := initializeRouter()
	router.Run(":8080")
}
