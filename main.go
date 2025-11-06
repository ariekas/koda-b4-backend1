package main

import (
	"crud/controllers"
	"crud/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "crud/docs" 
)

// @BasePath /
func main() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20
	r.Use(controllers.CrosMiddelware)
	r.Use(controllers.AllowPrelight)
	routes.AuthRouter(r)
	routes.UsersRouter(r)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
