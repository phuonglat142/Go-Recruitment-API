package main

import (
	"crud/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Đọc cấu hình từ file .env hoặc biến môi trường'
	cfg := config.Load()

	//Ở production tắt log debug của gin
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
