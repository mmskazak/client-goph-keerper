package commands

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var saveFileCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		filePath, _ := cmd.Flags().GetString("file")

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Добавляем заголовок
		writer.WriteField("title", title)
		writer.WriteField("description", description)

		// Открываем файл для чтения
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("ошибка открытия файла: %v", err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile("file", file.Name())
		if err != nil {
			return fmt.Errorf("ошибка создания формы файла: %v", err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return fmt.Errorf("ошибка копирования файла: %v", err)
		}

		writer.Close()

		// Отправляем запрос
		req, err := http.NewRequest("POST", "http://localhost:8080/file/save", body)
		if err != nil {
			return err
		}
		// Получаем токен из базы данных
		token, err := getTokenFromDB()
		if err != nil {
			return fmt.Errorf("ошибка при получении токена: %v", err)
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Authorization", token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		fmt.Printf("Response: %v\n", resp.Status)
		return nil
	},
}

func InitSaveFileCmd() *cobra.Command {
	saveFileCmd.Flags().String("title", "", "Title of the file")
	saveFileCmd.Flags().String("description", "", "Description of the file")
	saveFileCmd.Flags().String("file", "", "Path to the file")
	saveFileCmd.MarkFlagRequired("title")
	saveFileCmd.MarkFlagRequired("file")
	return saveFileCmd
}
