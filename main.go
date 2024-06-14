package main

import (
	"github.com/JoaoFerrareis02/REST-API-GO/db"
	"github.com/JoaoFerrareis02/REST-API-GO/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080

}
