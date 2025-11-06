package routes

import (
	"crud/controllers"
	"crud/middelware"

	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.Engine) {
	r.GET("/users", middelware.VerifToken(), func(ctx *gin.Context) {
		controllers.GetAll(ctx)
	})

	r.GET("/users/:id",middelware.VerifToken(), func(ctx *gin.Context) {
		controllers.GetById(ctx)
	})

	r.POST("/users",middelware.VerifToken(), func(ctx *gin.Context) {
		controllers.Create(ctx)
	})

	r.PATCH("/users/:id",middelware.VerifToken(), func(ctx *gin.Context) {
		controllers.Edit(ctx)
	})

	r.DELETE("/users/:id",middelware.VerifToken(), func(ctx *gin.Context) {
		controllers.Delete(ctx)
	})

	r.PATCH("/update/profile/:id",middelware.VerifToken(), func(ctx *gin.Context) {
		controllers.UploadProfile(ctx)
	})
}
