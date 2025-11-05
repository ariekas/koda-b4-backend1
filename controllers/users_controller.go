package controllers

import (
	"crud/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

// GetAll godoc
// @Summary Get all users
// @Description Mengambil semua data user
// @Tags Users
// @Produce json
// @Success 200 {object} models.Response
// @Router /users [get]
func GetAll(ctx *gin.Context){
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "success!",
		Data:    models.Users,
	})
}

// GetById godoc
// @Summary Get user by ID
// @Description Mengambil data user berdasarkan ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [get]
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

// Create godoc
// @Summary Create new user
// @Description Membuat user baru
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users [post]
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

// Edit godoc
// @Summary Update user by ID
// @Description Mengupdate data user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User Data"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [put]
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

// Delete godoc
// @Summary Delete user by ID
// @Description Menghapus user berdasarkan ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [delete]
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
