package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CrosMiddelware(ctx *gin.Context) {
	godotenv.Load()
	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ORIGIN_URL"))
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func AllowPrelight(ctx *gin.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(200)
	} else {
		ctx.Next()
	}
}
