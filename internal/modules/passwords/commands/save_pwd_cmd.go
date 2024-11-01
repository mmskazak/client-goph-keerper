package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	_ "github.com/glebarez/sqlite" // Импорт драйвера SQLite
	"github.com/spf13/cobra"
	"net/http"
)

// SetSavePasswordCmd команда сохранения пароля
func SetSavePasswordCmd(s *storage.Storage) (*cobra.Command, error) {
	savePwdCmd := &cobra.Command{
		Use:   "save",
		Short: "Save a password",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Получаем значения флагов
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")
			login, _ := cmd.Flags().GetString("login")
			password, _ := cmd.Flags().GetString("password")

			// Формируем JSON-данные для отправки
			data := map[string]interface{}{
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
			req, err := http.NewRequest("POST", s.ServerURL+"/pwd/save", bytes.NewBuffer(body))
			if err != nil {
				return err
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", s.Token)

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

	savePwdCmd.Flags().String("title", "", "Title for the password entry")
	savePwdCmd.Flags().String("description", "", "Description for the password entry")
	savePwdCmd.Flags().String("login", "", "Login for the password entry")
	savePwdCmd.Flags().String("password", "", "Password for the password entry")

	err := savePwdCmd.MarkFlagRequired("title")
	if err != nil {
		return nil, fmt.Errorf("error setting requeired title: %v", err)
	}
	err = savePwdCmd.MarkFlagRequired("login")
	if err != nil {
		return nil, fmt.Errorf("error setting requeired login: %v", err)
	}
	err = savePwdCmd.MarkFlagRequired("password")
	if err != nil {
		return nil, fmt.Errorf("error setting requeired password: %v", err)
	}

	return savePwdCmd, nil
}
