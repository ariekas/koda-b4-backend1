package controllers

import (
	"crud/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
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
	argon := argon2.DefaultConfig()

	err := ctx.BindJSON(&newuser)
	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Invalid JSON body!",
		})
		return
	}

	encoded, err := argon.HashEncoded([]byte(newuser.Password))
	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "",
		})
	}

	newuser.Password = string(encoded)

	models.Users = append(models.Users, newuser)

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Success Create User!",
		Data:    []models.User{newuser},
	})
}

func Edit(ctx *gin.Context){
	id := ctx.Param("id")
	argon := argon2.DefaultConfig()
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
			if newuser.Name != "" {
				models.Users[i].Name = newuser.Name
			}
			if newuser.Email != "" {
				models.Users[i].Email = newuser.Email
			}
			if newuser.Password != "" {
				encoded, err := argon.HashEncoded([]byte(newuser.Password))
				if err != nil {
					ctx.JSON(400, models.Response{
						Success: false,
						Message: "Error hashing password",
					})
					return
				}
				models.Users[i].Password = string(encoded)
			}
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "Success updated user!",
				Data:    []models.User{models.Users[i]},
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
