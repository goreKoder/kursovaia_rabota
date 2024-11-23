package main

import (
	"image-processing-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Роут для загрузки изображения
	r.POST("/upload", handlers.UploadImage)

	// Запуск сервера
	r.Run(":8080") // API будет доступен на localhost:8080
}
