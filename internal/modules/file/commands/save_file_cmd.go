package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

// SetSaveFileCmd создает команду для сохранения файла.
func SetSaveFileCmd(s *storage.Storage) (*cobra.Command, error) {
	saveFileCmd := &cobra.Command{
		Use:   "save",
		Short: "Save a file",
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath, err := cmd.Flags().GetString("file")
			if err != nil {
				return fmt.Errorf("ошибка получения пути к файлу: %w", err)
			}

			fileName := filepath.Base(filePath) // Извлекаем имя файла из пути

			// Создаем буфер для хранения тела запроса
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			// Открываем файл для чтения
			file, err := os.Open(filePath)
			if err != nil {
				return fmt.Errorf("ошибка открытия файла: %w", err)
			}
			defer func() {
				if err := file.Close(); err != nil {
					log.Printf("ошибка закрытия файла: %v", err)
				}
			}()

			// Создаем часть формы для загружаемого файла
			part, err := writer.CreateFormFile("file", fileName) // Используем fileName вместо file.Name()
			if err != nil {
				return fmt.Errorf("ошибка создания формы файла: %v", err)
			}

			// Копируем содержимое файла в часть формы
			if _, err := io.Copy(part, file); err != nil {
				return fmt.Errorf("ошибка копирования файла: %v", err)
			}

			// Закрываем writer после добавления всех частей
			if err := writer.Close(); err != nil {
				return fmt.Errorf("ошибка закрытия writer: %v", err)
			}

			// Создаем HTTP-запрос
			req, err := http.NewRequest(http.MethodPost, path.Join(s.ServerURL, "file", "save"), body) // Используем path.Join
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %v", err)
			}

			// Устанавливаем заголовки запроса
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Authorization", s.Token)

			// Отправляем запрос
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %v", err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Printf("error closing resp.Body: %v", err)
				}
			}(resp.Body)

			// Проверяем статус ответа
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("ошибка при сохранении файла, статус: %v", resp.Status)
			}

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	// Добавление флага для указания файла
	saveFileCmd.Flags().String("file", "", "Путь к файлу для сохранения")
	err := saveFileCmd.MarkFlagRequired("file")
	if err != nil {
		return nil, fmt.Errorf("error mark file flag as required: %w", err)
	}

	return saveFileCmd, nil
}
