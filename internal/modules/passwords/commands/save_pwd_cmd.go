package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	_ "github.com/glebarez/sqlite" // Импорт драйвера SQLite
	"github.com/spf13/cobra"
)

// SetSavePasswordCmd команда сохранения пароля.
func SetSavePasswordCmd(s *storage.Storage) (*cobra.Command, error) {
	savePwdCmd := &cobra.Command{
		Use:   "save",
		Short: "Save a password",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Получаем значения флагов
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				return fmt.Errorf("get title flag: %w", err)
			}
			login, err := cmd.Flags().GetString("login")
			if err != nil {
				return fmt.Errorf("get login flag: %w", err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				return fmt.Errorf("get password flag: %w", err)
			}

			// Формируем JSON-данные для отправки
			data := map[string]interface{}{
				"title": title,
				"credentials": map[string]string{
					"login":    login,
					"password": password,
				},
			}

			body, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("ошибка кодирования JSON: %w", err)
			}

			// Создаем и отправляем запрос
			reqURL := path.Join(s.ServerURL, Pwd, "save")
			req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(body))
			if err != nil {
				return fmt.Errorf("error saving password: %w", err)
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", s.Token)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("error saving password: %w", err)
			}
			defer resp.Body.Close() //nolint:errcheck //опустим здесь проверку

			fmt.Printf(Response, resp.Status)
			return nil
		},
	}

	savePwdCmd.Flags().String("title", "", "Title for the password entry")
	savePwdCmd.Flags().String("login", "", "Login for the password entry")
	savePwdCmd.Flags().String("password", "", "Password for the password entry")

	err := savePwdCmd.MarkFlagRequired("title")
	if err != nil {
		return nil, fmt.Errorf("error setting requeired title: %w", err)
	}
	err = savePwdCmd.MarkFlagRequired("login")
	if err != nil {
		return nil, fmt.Errorf("error setting requeired login: %w", err)
	}
	err = savePwdCmd.MarkFlagRequired("password")
	if err != nil {
		return nil, fmt.Errorf("error setting requeired password: %w", err)
	}

	return savePwdCmd, nil
}
