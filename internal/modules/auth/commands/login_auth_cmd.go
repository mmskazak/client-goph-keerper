package commands

import (
	"bytes"
	"client-goph-keerper/internal/storage"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

// LoginCommand инициализирует команду для входа пользователя
func LoginCommand(s *storage.Storage) (*cobra.Command, error) {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Log in a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Проверяем, есть ли уже сохраненный токен
			isHaveToken, err := checkTokenExists()
			if err != nil {
				return fmt.Errorf("checkTokenExists: %w", err)
			}
			if isHaveToken {
				return fmt.Errorf("пользователь уже авторизован")
			}

			// Получаем логин и пароль из флагов
			login, _ := cmd.Flags().GetString("login")
			password, _ := cmd.Flags().GetString("password")

			// Формируем данные для JSON-запроса
			data := map[string]string{
				"login":    login,
				"password": password,
			}

			body, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("ошибка кодирования JSON: %v", err)
			}

			// Выполняем HTTP-запрос для входа
			req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(body))
			if err != nil {
				return err
			}

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// Читаем ответ от сервера
			all, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Println(string(all))

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("Response: %v\n", resp.Status)
				return nil
			}

			// Извлекаем токен из заголовка ответа
			token := resp.Header.Get("Authorization")
			if token == "" {
				return fmt.Errorf("токен не найден в заголовке")
			}

			// Сохраняем токен в базе данных
			_, err = s.DataBase.Exec(`
				INSERT OR REPLACE INTO app_params (key, value)
				VALUES (?, ?)`, "jwt_token", token)
			if err != nil {
				return fmt.Errorf("failed to save JWT token: %v", err)
			}
			fmt.Println("JWT токен успешно сохранен:", token)

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	// Добавляем обязательные флаги login и password
	loginCmd.Flags().String("login", "", "Login for the user")
	loginCmd.Flags().String("password", "", "Password for the user")
	err := loginCmd.MarkFlagRequired("login")
	if err != nil {
		return nil, fmt.Errorf("failed to mark `login` flag as required: %v", err)
	}
	err = loginCmd.MarkFlagRequired("password")
	if err != nil {
		return nil, fmt.Errorf("failed to mark `password` flag as required: %v", err)
	}

	return loginCmd, nil
}
