package controllers

import (
	"crud/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context){
	var login models.User
		login.Email = ctx.PostForm("email")
		login.Password = ctx.PostForm("password")

		for _, user := range models.Users {
			if user.Email == login.Email && user.Password == login.Password {
				ctx.JSON(200, models.Response{
					Success: true,
					Message: "Login berhasil!",
					Data:    []models.User{user},
				})
				return
			}else{
				ctx.JSON(404, models.Response{
					Success: false,
					Message: "Email atau password salah!",
				})
			}
		}
}

func AuthRegister(ctx *gin.Context){
	var newuser models.User
		newuser.Id = len(models.Users) + 1
		newuser.Name = ctx.PostForm("name")
		newuser.Email = ctx.PostForm("email")
		newuser.Password = ctx.PostForm("password")

		if !strings.Contains(newuser.Email, "@") {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Wrong email type",
			})
			return
		}

		if len(newuser.Password) < 8 {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Password Much 8 carakter",
			})
			return
		}
		models.Users = append(models.Users, newuser)

		ctx.JSON(200, models.Response{
			Success: true,
			Message: "Success Create User!",
			Data:    []models.User{newuser},
		})
}