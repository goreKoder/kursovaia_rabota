package main //		go run main.go

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// Открываем файл
	filePath := "image.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Создаем буфер для тела запроса
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем файл в тело запроса
	part, err := writer.CreateFormFile("image", "image.jpg")
	if err != nil {
		fmt.Println("Ошибка создания файла в запросе:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Ошибка записи файла в запрос:", err)
		return
	}

	// Завершаем формирование запроса
	writer.Close()

	// Отправляем запрос
	url := "http://localhost:8080/upload"
	req, err := http.NewRequest("POST", url, body) //вроде как оно расширяет возможности отправки запроса, но мне похер, я и так ели-ели понимаю как работает мой код
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer resp.Body.Close()
	defer req.Body.Close()

	// Читаем ответ
	// responseBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Ошибка чтения ответа:", err)
	// 	return
	// }
	// fmt.Println("Ответ сервера:", string(responseBody))

	// Открываем файл для записи на диске
	outFile, err := os.Create("downloaded_image.jpg")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outFile.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

}
