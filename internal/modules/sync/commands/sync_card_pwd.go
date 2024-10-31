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

type CardEntry struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	PinCode     int    `json:"pincode"`
	CVV         int    `json:"cvv"`
	Expire      string `json:"expire"`
}

// Команда для получения всех карт
var allCardsCmd = &cobra.Command{
	Use:   "cards",
	Short: "Synchronize the cards",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Создаем запрос
		req, err := http.NewRequest("POST", "http://localhost:8080/cards/all", nil)
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
		var entries []CardEntry
		if err := json.Unmarshal(responseData, &entries); err != nil {
			return fmt.Errorf("ошибка разбора JSON: %v", err)
		}

		// Сохраняем записи в базу данных
		if err := saveCardsToDB(entries); err != nil {
			return fmt.Errorf("ошибка при сохранении карт в базу данных: %v", err)
		}

		fmt.Println("Все карты успешно сохранены в локальную базу данных.")
		return nil
	},
}

// Функция для сохранения карт в базу данных
func saveCardsToDB(entries []CardEntry) error {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Вставляем каждую карту в таблицу
	insertQuery := `INSERT INTO cards (title, description, number, pincode, cvv, expire) VALUES (?, ?, ?, ?, ?, ?)`
	for _, entry := range entries {
		if _, err := db.Exec(insertQuery, entry.Title, entry.Description, entry.Number, entry.PinCode, entry.CVV, entry.Expire); err != nil {
			return fmt.Errorf("ошибка вставки карты в базу данных: %v", err)
		}
	}

	return nil
}

func InitSyncAllCardsCmd() *cobra.Command {
	return allCardsCmd
}
