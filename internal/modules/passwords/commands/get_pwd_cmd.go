package commands

import (
	"client-goph-keerper/internal/storage"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/spf13/cobra"
)

// SetGetPasswordCmd создает команду получения пароля по ID.
func SetGetPasswordCmd(s *storage.Storage) (*cobra.Command, error) {
	getPwdCmd := &cobra.Command{
		Use:   "get",
		Short: "Get a password by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Получаем значение флага
			pwdID, _ := cmd.Flags().GetString("pwd_id")

			// Формируем URL запроса
			url := path.Join(s.ServerURL, "pwd", "get", pwdID)

			// Создаем запрос
			req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %w", err)
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", s.Token)

			// Отправляем запрос
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %w", err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Printf("error closing body: %v", err)
				}
			}(resp.Body)

			// Чтение ответа
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("ошибка чтения ответа: %w", err)
			}

			fmt.Printf("Response: %v\n", resp.Status)
			fmt.Printf("Data: %s\n", data)
			return nil
		},
	}

	// Определяем флаги
	getPwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	// Устанавливаем обязательные флаги
	err := getPwdCmd.MarkFlagRequired("pwd_id")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'pwd_id': %w", err)
	}

	return getPwdCmd, nil
}
