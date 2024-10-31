package commands

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

// Структура для хранения информации о пароле
type PasswordEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}

// Команда для получения всех паролей
var allPwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Synchronize the password",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Создаем запрос
		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/all", nil)
		if err != nil {
			return err
		}

		// Получаем токен из базы данных
		token, err := getTokenFromDB()
		if err != nil {
			return fmt.Errorf("ошибка при получении токена: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Чтение тела ответа
		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %v", err)
		}

		// Парсим JSON-ответ
		var entries []PasswordEntry
		if err := json.Unmarshal(responseData, &entries); err != nil {
			return fmt.Errorf("ошибка разбора JSON: %v", err)
		}

		// Сохраняем записи в базу данных
		if err := savePasswordsToDB(entries); err != nil {
			return fmt.Errorf("ошибка при сохранении паролей в базу данных: %v", err)
		}

		fmt.Println("Все пароли успешно сохранены в локальную базу данных.")
		return nil
	},
}

// Функция для сохранения паролей в базу данных
func savePasswordsToDB(entries []PasswordEntry) error {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Вставляем каждый пароль в таблицу
	insertQuery := `INSERT INTO passwords (title, description, login, password) VALUES (?, ?, ?, ?)`
	for _, entry := range entries {
		if _, err := db.Exec(insertQuery, entry.Title, entry.Description, entry.Login, entry.Password); err != nil {
			return fmt.Errorf("ошибка вставки пароля в базу данных: %v", err)
		}
	}

	return nil
}

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

func InitSyncAllPwdCmd() *cobra.Command {
	return allPwdCmd
}
