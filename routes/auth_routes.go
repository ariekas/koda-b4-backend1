package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/register", func(ctx *gin.Context) {
		controllers.AuthRegister(ctx)
	})

	r.POST("/login", func(ctx *gin.Context) {
		controllers.AuthLogin(ctx)
	})
}
