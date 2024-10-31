package commands

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	RunE: func(cmd *cobra.Command, args []string) error {

		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")

		data := map[string]string{
			"login":    login,
			"password": password,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

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
		all, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		fmt.Println(string(all))

		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		}

		token := resp.Header.Get("Authorization")
		if token == "" {
			return fmt.Errorf("токен не найден в заголовке")
		}
		err = saveTokenToDB(token)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("JWT токен:", token)

		fmt.Printf("Response: %v\n", resp.Status)
		return nil
	},
}

func InitRegisterCmd() *cobra.Command {
	registerCmd.Flags().String("login", "", "Login for the new user")
	registerCmd.Flags().String("password", "", "Password for the new user")
	registerCmd.MarkFlagRequired("login")
	registerCmd.MarkFlagRequired("password")
	return registerCmd
}

func saveTokenToDB(jwt string) error {
	// Подключаемся к базе данных с драйвером glebarez/sqlite
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Вставляем токен в таблицу
	insertQuery := `INSERT INTO users (jwt) VALUES (?)`
	if _, err := db.Exec(insertQuery, jwt); err != nil {
		return fmt.Errorf("ошибка вставки токена в базу данных: %v", err)
	}

	fmt.Println("Токен успешно сохранен в базу данных.")
	return nil
}

func checkTokenExists() (bool, error) {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return false, fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Проверяем наличие токена
	query := `SELECT EXISTS(SELECT * FROM users WHERE 1)`
	var exists bool
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}

	return exists, nil
}
