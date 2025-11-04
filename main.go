package main

import (
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	routes.AuthRouter(r)
	routes.UsersRouter(r)

	r.Run(":8080")
}
