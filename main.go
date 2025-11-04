package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var users = []User{
	// {Id: 1, Name: "ari eka", Email: "ari@gmail.com", Password: "1234"},
	// {Id: 2, Name: "yanto", Email: "ynt@gmail.com", Password: "12"},
	// {Id: 3, Name: "paw paw", Email: "pkw@gmail.com", Password: "ppack"},
	// {Id: 4, Name: "knakri", Email: "aw@gmail.com", Password: "aw123"},
}

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "success!",
			Data:    users,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, user := range users {
			if fmt.Sprint(user.Id) == id {
				ctx.JSON(200, Response{
					Success: true,
					Message: "Getting data user!",
					Data:    []User{user},
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "User not found",
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		var newuser User

		err := ctx.BindJSON(&newuser)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid JSON body!",
			})
			return
		}

		users = append(users, newuser)

		ctx.JSON(200, Response{
			Success: true,
			Message: "Success Create User!",
			Data:    []User{newuser},
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var newuser User

		err := ctx.BindJSON(&newuser)
		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Invalid JSON body!",
			})
			return
		}

		for i, user := range users {
			if fmt.Sprint(user.Id) == id {
				users = append(users[:i], []User{newuser}...)
				ctx.JSON(200, Response{
					Success: true,
					Message: "Succes updated user!",
					Data:    []User{newuser},
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "User not found",
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i, user := range users {
			if fmt.Sprint(user.Id) == id {
				users = append(users[:i], users[i+1:]...)

				ctx.JSON(200, Response{
					Success: true,
					Message: "Deleting data!",
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "User not found",
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
		var newuser User
		newuser.Id = len(users) + 1
		newuser.Name = ctx.PostForm("name")
		newuser.Email = ctx.PostForm("email")
		newuser.Password = ctx.PostForm("password")

		if !strings.Contains(newuser.Email, "@") {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Wrong email type",
			})
			return
		}

		if len(newuser.Password) < 8 {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Password Much 8 carakter",
			})
			return
		}
		users = append(users, newuser)

		ctx.JSON(200, Response{
			Success: true,
			Message: "Success Create User!",
			Data:    []User{newuser},
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var login User
		login.Email = ctx.PostForm("email")
		login.Password = ctx.PostForm("password")

		for _, user := range users {
			if user.Email == login.Email && user.Password == login.Password {
				ctx.JSON(200, Response{
					Success: true,
					Message: "Login berhasil!",
					Data:    []User{user},
				})
				return
			}else{
				ctx.JSON(404, Response{
					Success: false,
					Message: "Email atau password salah!",
				})
			}
		}
	})

	r.Run(":8080")
}
