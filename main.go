package main

import (
	"github.com/edu2801/first-golang-crud/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run(":8001")
}
