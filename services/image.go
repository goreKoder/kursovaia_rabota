package services

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func ProcessImage(filePath string) (string, error) {
	// Открытие исходного изображения
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
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
	processedPath := "./uploads/processed_image.jpg"
	outputFile, err := os.Create(processedPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, grayImg, nil)
	if err != nil {
		return "", err
	}

	return processedPath, nil
}
