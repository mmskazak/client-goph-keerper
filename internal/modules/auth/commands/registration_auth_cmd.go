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

// RegisterCommand создаёт команду регистрации пользователя
func RegisterCommand(s *storage.Storage) (*cobra.Command, error) {
	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Получаем логин и пароль из флагов
			login, _ := cmd.Flags().GetString("login")
			password, _ := cmd.Flags().GetString("password")

			// Создаем JSON объект для передачи на сервер
			data := map[string]string{
				"login":    login,
				"password": password,
			}

			body, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("ошибка кодирования JSON: %v", err)
			}

			// Создаём HTTP запрос для регистрации
			req, err := http.NewRequest("POST", "http://localhost:8080/registration", bytes.NewBuffer(body))
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

			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("Response: %v\n", resp.Status)
				return nil
			}

			// Получаем токен из заголовка ответа
			token := resp.Header.Get("Authorization")
			if token == "" {
				return fmt.Errorf("токен не найден в заголовке")
			}

			// Сохраняем токен в базе данных
			err = saveTokenToDB(s, token)
			if err != nil {
				return fmt.Errorf("ошибка сохранения токена: %v", err)
			}
			fmt.Println("JWT токен:", token)

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	registerCmd.Flags().String("login", "", "Login for the new user")
	registerCmd.Flags().String("password", "", "Password for the new user")
	registerCmd.MarkFlagRequired("login")
	registerCmd.MarkFlagRequired("password")

	return registerCmd, nil
}

// saveTokenToDB сохраняет токен в базе данных, используя переданный объект *storage.Storage
func saveTokenToDB(s *storage.Storage, jwt string) error {
	insertQuery := `INSERT INTO users (jwt) VALUES (?)`
	if _, err := s.DataBase.Exec(insertQuery, jwt); err != nil {
		return fmt.Errorf("ошибка вставки токена в базу данных: %v", err)
	}

	fmt.Println("Токен успешно сохранен в базу данных.")
	return nil
}

// checkTokenExists проверяет наличие токена в базе данных
func checkTokenExists(s *storage.Storage) (bool, error) {
	query := `SELECT EXISTS(SELECT * FROM users WHERE 1)`
	var exists bool
	err := s.DataBase.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}

	return exists, nil
}
