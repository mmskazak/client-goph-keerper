package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// SetSaveFileCmd создает команду для сохранения файла
func SetSaveFileCmd(s *storage.Storage) (*cobra.Command, error) {
	saveFileCmd := &cobra.Command{
		Use:   "save",
		Short: "Save a file",
		RunE: func(cmd *cobra.Command, args []string) error {
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")
			filePath, _ := cmd.Flags().GetString("file")

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			// Добавляем заголовки
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
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/file/save", s.ServerURL), body)
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %v", err)
			}

			// Получаем токен из базы данных
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", s.Token)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %v", err)
			}
			defer resp.Body.Close()

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	saveFileCmd.Flags().String("title", "", "Title of the file")
	saveFileCmd.Flags().String("description", "", "Description of the file")
	saveFileCmd.Flags().String("file", "", "Path to the file")
	err := saveFileCmd.MarkFlagRequired("title")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'title': %v", err)
	}
	err = saveFileCmd.MarkFlagRequired("file")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'file': %v", err)
	}

	return saveFileCmd, nil
}
