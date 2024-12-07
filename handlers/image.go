package handlers

import (
	"fmt"
	"image-processing-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	// Извлекаем параметр id
	id := c.Param("id")
	fmt.Println("id = ", id)
	// Получение файла из запроса
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл"})
		fmt.Println("error: Не удалось получить файл")
		return
	}

	// Сохранение файла на диск (опционально)
	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл"})
		fmt.Println("error: Не удалось сохранить файл")
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	switch i {
	case 1: //		обрезка
		processedPath, err := services.CropImage(filePath, 100, 100, 200, 200)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки изображения"})
			fmt.Println("error: Ошибка обработки изображения")
			return
		}
		c.File(processedPath)
	case 2: //		изменение размера
		processedPath, err := services.ResizeImage(filePath, 150, 150)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки изображения"})
			fmt.Println("error: Ошибка обработки изображения")
			return
		}
		c.File(processedPath)
	case 3:
		// Смена цвета
		processedPath, err := services.ProcessImage(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки изображения"})
			fmt.Println("error: Ошибка обработки изображения")
			return
		}
		c.File(processedPath)
	}
}
