package commands

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/glebarez/sqlite" // Импорт драйвера SQLite
	"github.com/spf13/cobra"
	"net/http"
)

// Функция для получения токена из базы данных
func getTokenFromDB() (string, error) {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Извлекаем токен из таблицы
	var token string
	query := `SELECT jwt FROM users LIMIT 1`
	err = db.QueryRow(query).Scan(&token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("токен не найден в базе данных")
		}
		return "", fmt.Errorf("ошибка получения токена: %v", err)
	}

	return token, nil
}

var savePwdCmd = &cobra.Command{
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
		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/save", bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		// Получаем токен из базы данных
		token, err := getTokenFromDB()
		if err != nil {
			return fmt.Errorf("ошибка при получении токена: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)

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

func InitSavePwdCmd() *cobra.Command {
	savePwdCmd.Flags().String("title", "", "Title for the password entry")
	savePwdCmd.Flags().String("description", "", "Description for the password entry")
	savePwdCmd.Flags().String("login", "", "Login for the password entry")
	savePwdCmd.Flags().String("password", "", "Password for the password entry")
	savePwdCmd.MarkFlagRequired("title")
	savePwdCmd.MarkFlagRequired("login")
	savePwdCmd.MarkFlagRequired("password")
	return savePwdCmd
}
