package commands

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type TextEntry struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Text        string `json:"text"`
}

// Команда для получения всех текстов
var allTextCmd = &cobra.Command{
	Use:   "texts",
	Short: "Synchronize the texts",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Создаем запрос
		req, err := http.NewRequest("POST", "http://localhost:8080/texts/all", nil)
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

		// Чтение тела ответа
		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %v", err)
		}
		fmt.Println("Тело ответа")
		fmt.Println(string(responseData))

		// Парсим JSON-ответ
		var entries []TextEntry
		if err := json.Unmarshal(responseData, &entries); err != nil {
			return fmt.Errorf("ошибка разбора JSON: %v", err)
		}

		// Сохраняем записи в базу данных
		if err := saveTextsToDB(entries); err != nil {
			return fmt.Errorf("ошибка при сохранении текстов в базу данных: %v", err)
		}

		fmt.Println("Все тексты успешно сохранены в локальную базу данных.")
		return nil
	},
}

// Функция для сохранения текстов в базу данных
func saveTextsToDB(entries []TextEntry) error {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Вставляем каждый текст в таблицу
	insertQuery := `INSERT INTO texts (title, description, text) VALUES (?, ?, ?)`
	for _, entry := range entries {
		if _, err := db.Exec(insertQuery, entry.Title, entry.Description, entry.Text); err != nil {
			return fmt.Errorf("ошибка вставки текста в базу данных: %v", err)
		}
	}

	return nil
}
