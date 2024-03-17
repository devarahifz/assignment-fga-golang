package main

import (
	"assignment3/routers"
	"fmt"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)

	fmt.Println("Server is running on port", PORT)
}
