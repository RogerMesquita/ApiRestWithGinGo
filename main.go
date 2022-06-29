package main

import (
	"ApiRestWithGinGo/database"
	"ApiRestWithGinGo/routes"
)

func main() {
	database.ConnectBd()

	routes.HandleRequests()

}
