package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

// SetDeletePasswordCmd создает команду удаления пароля по ID
func SetDeletePasswordCmd(s *storage.Storage) (*cobra.Command, error) {
	deletePwdCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a password by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Получаем значение флага
			pwdID, _ := cmd.Flags().GetString("pwd_id")

			// Формируем JSON-данные для отправки
			data := map[string]interface{}{
				"pwd_id": pwdID,
			}

			body, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("ошибка кодирования JSON: %v", err)
			}

			// Создаем запрос
			url := fmt.Sprintf("%s/pwd/delete", s.ServerURL)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	// Определяем флаг
	deletePwdCmd.Flags().String("pwd_id", "", "Password entry ID")

	// Устанавливаем обязательный флаг
	err := deletePwdCmd.MarkFlagRequired("pwd_id")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'pwd_id': %v", err)
	}

	return deletePwdCmd, nil
}
