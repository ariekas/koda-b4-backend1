package main

import (
	"crud/controllers"
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(controllers.CrosMiddelware)
	r.Use(controllers.AllowPrelight)
	routes.AuthRouter(r)
	routes.UsersRouter(r)

	r.Run(":8080")
}
