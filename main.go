package main

import (
	"image-processing-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode) // tckb 'nj yt chf,jnftn то я буду плакать
	r := gin.Default()
	// Роут для загрузки изображения
	r.POST("/upoad/:id", handlers.UploadImage)

	// Запуск сервера
	r.Run(":8080") // Слушаем порт 8080
}

//		go run main.go
