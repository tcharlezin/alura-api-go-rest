package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConnectToDatabase()

	fmt.Println("Iniciando servidor REST com GO!")
	routes.HandleRequest()
}
