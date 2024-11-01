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

type FileEntry struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PathToFile  string `json:"path_to_file"`
}

// Команда для получения всех файлов
var allFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Synchronize the files",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Создаем запрос
		req, err := http.NewRequest("POST", "http://localhost:8080/files/all", nil)
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
		var entries []FileEntry
		if err := json.Unmarshal(responseData, &entries); err != nil {
			return fmt.Errorf("ошибка разбора JSON: %v", err)
		}

		// Сохраняем записи в базу данных
		if err := saveFilesToDB(entries); err != nil {
			return fmt.Errorf("ошибка при сохранении файлов в базу данных: %v", err)
		}

		fmt.Println("Все файлы успешно сохранены в локальную базу данных.")
		return nil
	},
}

// Функция для сохранения файлов в базу данных
func saveFilesToDB(entries []FileEntry) error {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Вставляем каждый файл в таблицу
	insertQuery := `INSERT INTO files (title, description, path_to_file) VALUES (?, ?, ?)`
	for _, entry := range entries {
		if _, err := db.Exec(insertQuery, entry.Title, entry.Description, entry.PathToFile); err != nil {
			return fmt.Errorf("ошибка вставки файла в базу данных: %v", err)
		}
	}

	return nil
}

func InitSyncAllFilesCmd() *cobra.Command {
	return allFilesCmd
}
