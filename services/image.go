package services

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func ProcessImage(filePath string) (string, error) {
	// Открытие исходного изображения
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Ошибка при открытии файла: ", err)
		return "nil", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Ошибка при декодировании файла: ", err)
		return "nil", err
	}
	// Пример обработки: конвертация в черно-белый
	grayImg := image.NewGray(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	// Сохранение обработанного изображения
	outputFile, err := os.Create("./uploads/downloaded_image.jpg")
	if err != nil {
		log.Println("Ошибка при сохранении обработанного изображения: ", err)
		return "nil", err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, grayImg, nil)
	if err != nil {
		log.Println("Ошибка при кодировании обработанного изображения: ", err)
		return "nil", err
	}
	return "./uploads/downloaded_image.jpg", nil

}

// cropImage обрезает изображение до заданного прямоугольника
func CropImage(filePath string, x, y, w, h int) (string, error) {
	// Открытие исходного изображения
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Ошибка при открытии файла: ", err)
		return "nil", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Ошибка при декодировании файла: ", err)
		return "nil", err
	}
	rect := image.Rect(x, y, x+w, y+h)
	cropped := image.NewRGBA(rect)
	draw.Draw(cropped, rect, img, rect.Min, draw.Src)
	// Сохранение обработанного изображения
	outputFile, err := os.Create("./uploads/downloaded_image.jpg")
	if err != nil {
		log.Println("Ошибка при сохранении обработанного изображения: ", err)
		return "nil", err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, cropped, nil)
	if err != nil {
		log.Println("Ошибка при кодировании обработанного изображения: ", err)
		return "nil", err
	}
	return "./uploads/downloaded_image.jpg", nil
}

// resizeImage изменяет размер изображения, сохраняя пропорции
func ResizeImage(filePath string, width, height int) (string, error) {
	// Открытие исходного изображения
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Ошибка при открытии файла: ", err)
		return "nil", err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Ошибка при декодировании файла: ", err)
		return "nil", err
	}

	img_cover := resize.Resize(500, 500, img, resize.NearestNeighbor)
	// Сохранение обработанного изображения
	outputFile, err := os.Create("./uploads/downloaded_image.jpg")
	if err != nil {
		log.Println("Ошибка при сохранении обработанного изображения: ", err)
		return "nil", err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, img_cover, nil)
	if err != nil {
		log.Println("Ошибка при кодировании обработанного изображения: ", err)
		return "nil", err
	}
	return "./uploads/downloaded_image.jpg", nil
}
