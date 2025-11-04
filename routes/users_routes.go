package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.Engine) {
	r.GET("/users", func(ctx *gin.Context) {
		controllers.GetAll(ctx)
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		controllers.GetById(ctx)
	})

	r.POST("/users", func(ctx *gin.Context) {
		controllers.Create(ctx)
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		controllers.Edit(ctx)
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		controllers.Delete(ctx)
	})
}
