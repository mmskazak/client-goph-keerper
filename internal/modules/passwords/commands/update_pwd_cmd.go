package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

// SetUpdatePasswordCmd создает команду обновления пароля
func SetUpdatePasswordCmd(s *storage.Storage) (*cobra.Command, error) {
	updatePwdCmd := &cobra.Command{
		Use:   "update",
		Short: "Update a password by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Получаем значения флагов
			pwdID, _ := cmd.Flags().GetString("pwd_id")
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")
			login, _ := cmd.Flags().GetString("login")
			password, _ := cmd.Flags().GetString("password")

			// Формируем JSON-данные для отправки
			data := map[string]interface{}{
				"pwd_id":      pwdID,
				"title":       title,
				"description": description,
				"credentials": map[string]string{
					"login":    login,
					"password": password,
				},
			}

			body, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("ошибка кодирования JSON: %v", err)
			}

			// Создаем и отправляем запрос
			req, err := http.NewRequest("POST", s.ServerURL+"/pwd/update", bytes.NewBuffer(body))
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")
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

	updatePwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	updatePwdCmd.Flags().String("title", "", "Title for the password entry")
	updatePwdCmd.Flags().String("description", "", "Description for the password entry")
	updatePwdCmd.Flags().String("login", "", "Login for the password entry")
	updatePwdCmd.Flags().String("password", "", "Password for the password entry")

	err := updatePwdCmd.MarkFlagRequired("pwd_id")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'pwd_id': %v", err)
	}
	err = updatePwdCmd.MarkFlagRequired("title")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'title': %v", err)
	}
	err = updatePwdCmd.MarkFlagRequired("login")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'login': %v", err)
	}
	err = updatePwdCmd.MarkFlagRequired("password")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'password': %v", err)
	}

	return updatePwdCmd, nil
}
