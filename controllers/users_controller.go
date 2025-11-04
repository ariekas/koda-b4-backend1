package controllers

import (
	"crud/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context){
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "success!",
		Data:    models.Users,
	})
}

func GetById(ctx *gin.Context){
	id := ctx.Param("id")
	for _, user := range models.Users {
		if fmt.Sprint(user.Id) == id {
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Getting data user!",
				Data:    []models.User{user},
			})
			return
		}
	}

	ctx.JSON(404, models.Response{
		Success: false,
		Message: "User not found",
	})
}

func Create(ctx *gin.Context){
	var newuser models.User

	err := ctx.BindJSON(&newuser)
	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Invalid JSON body!",
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

func Edit(ctx *gin.Context){
	id := ctx.Param("id")
	var newuser models.User

	err := ctx.BindJSON(&newuser)
	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Invalid JSON body!",
		})
		return
	}

	for i, user := range models.Users {
		if fmt.Sprint(user.Id) == id {
			models.Users = append(models.Users[:i], []models.User{newuser}...)
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Succes updated user!",
				Data:    []models.User{newuser},
			})
			return
		}
	}

	ctx.JSON(404, models.Response{
		Success: false,
		Message: "User not found",
	})
}

func Delete(ctx *gin.Context){
	id := ctx.Param("id")
	for i, user := range models.Users {
		if fmt.Sprint(user.Id) == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)

			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Deleting data!",
			})
			return
		}
	}

	ctx.JSON(404, models.Response{
		Success: false,
		Message: "User not found",
	})
}
