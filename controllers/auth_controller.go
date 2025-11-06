package controllers

import (
	"crud/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

// AuthLogin godoc
// @Summary Login user
// @Description Melakukan login user menggunakan email dan password
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "Email user"
// @Param password formData string true "Password user"
// @Success 200 {object} models.Response "Login berhasil"
// @Failure 400 {object} models.Response "Email atau password salah"
// @Failure 404 {object} models.Response "User tidak ditemukan"
// @Router /login [post]
func AuthLogin(ctx *gin.Context){
	var login models.User
		login.Email = ctx.PostForm("email")
		login.Password = ctx.PostForm("password")

		for _, user := range models.Users {
			if user.Email == login.Email {
				ok, err := argon2.VerifyEncoded([]byte(login.Password), []byte(user.Password))
				if err != nil || !ok {
					ctx.JSON(400, models.Response{
						Success: false,
						Message: "Wrong email or password",
					})
					return
				}
	
				ctx.JSON(200, models.Response{
					Success: true,
					Message: "Success Login!",
					Data:    []models.User{user},
				})
				return
			}
		}
	
		ctx.JSON(404, models.Response{
			Success: false,
			Message: "User Not Found!",
		})
}

// AuthRegister godoc
// @Summary Register user baru
// @Description Mendaftarkan user baru menggunakan name, email, dan password
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nama lengkap"
// @Param email formData string true "Email user"
// @Param password formData string true "Password minimal 8 karakter"
// @Success 200 {object} models.Response "Registrasi berhasil"
// @Failure 400 {object} models.Response "Input tidak valid"
// @Router /register [post]
func AuthRegister(ctx *gin.Context){
	argon := argon2.DefaultConfig()
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