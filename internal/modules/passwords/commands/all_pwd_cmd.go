package commands

import (
	"client-goph-keerper/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

// SetAllPasswordsCmd создает команду для получения всех паролей пользователя
func SetAllPasswordsCmd(s *storage.Storage) (*cobra.Command, error) {
	allPwdCmd := &cobra.Command{
		Use:   "all",
		Short: "List all passwords for a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Формируем URL для запроса
			url := fmt.Sprintf("%s/pwd/all", s.ServerURL)

			// Создаем запрос
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %v", err)
			}

			// Устанавливаем заголовки
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", s.Token)

			// Отправляем запрос
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %v", err)
			}
			defer resp.Body.Close()

			// Чтение тела ответа
			responseData, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("ошибка чтения ответа: %v", err)
			}

			fmt.Printf("Status: %v\n", resp.Status)
			fmt.Printf("Response: %s\n", responseData)
			return nil
		},
	}

	return allPwdCmd, nil
}
