package main

import (
	"assignment2/database"
	"assignment2/routers"
	"fmt"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	fmt.Println("Server is running on port", PORT)

	routers.StartServer().Run(PORT)
}
