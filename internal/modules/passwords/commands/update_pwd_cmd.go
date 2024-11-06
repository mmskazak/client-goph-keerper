package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// SetUpdatePasswordCmd создает команду обновления пароля.
func SetUpdatePasswordCmd(s *storage.Storage) (*cobra.Command, error) {
	updatePwdCmd := &cobra.Command{
		Use:   "update",
		Short: "Update a password by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Получаем значения флагов
			pwdID, err := cmd.Flags().GetString("pwd_id")
			if err != nil {
				return fmt.Errorf("error getting pwd_id from flags: %w", err)
			}
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				return fmt.Errorf("error getting title from flags: %w", err)
			}
			login, err := cmd.Flags().GetString("login")
			if err != nil {
				return fmt.Errorf("error getting login from flags: %w", err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				return fmt.Errorf("error getting password from flags: %w", err)
			}

			// Формируем JSON-данные для отправки
			data := map[string]interface{}{
				"pwd_id": pwdID,
				"title":  title,
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
			req, err := http.NewRequest(http.MethodPost, s.ServerURL+"/pwd/update", bytes.NewBuffer(body))
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
	updatePwdCmd.Flags().String("login", "", "Login for the password entry")
	updatePwdCmd.Flags().String("password", "", "Password for the password entry")

	err := updatePwdCmd.MarkFlagRequired("pwd_id")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'pwd_id': %w", err)
	}
	err = updatePwdCmd.MarkFlagRequired("title")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'title': %w", err)
	}
	err = updatePwdCmd.MarkFlagRequired("login")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'login': %w", err)
	}
	err = updatePwdCmd.MarkFlagRequired("password")
	if err != nil {
		return nil, fmt.Errorf("error setting required flag 'password': %w", err)
	}

	return updatePwdCmd, nil
}
