package handlers

import (
	"image-processing-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	// Получение файла из запроса
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл"})
		return
	}

	// Сохранение файла на диск (опционально)
	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл"})
		return
	}

	// Обработка изображения
	processedPath, err := services.ProcessImage(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки изображения"})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message":        "Изображение обработано успешно",
		"processed_file": processedPath,
	})
}
