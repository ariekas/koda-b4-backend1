package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data  []User `json:"data"`
}

var users = []User{
	{Id: 1, Name: "ari eka", Email: "ari@gmail.com", Password: "1234"},
	{Id: 2, Name: "yanto", Email: "ynt@gmail.com", Password: "12"},
	{Id: 3, Name: "paw paw", Email: "pkw@gmail.com", Password: "ppack"},
	{Id: 4, Name: "knakri", Email: "aw@gmail.com", Password: "aw123"},
}


func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "success!",
			Data: users,
		})
	})

	r.Run(":8080")
}
