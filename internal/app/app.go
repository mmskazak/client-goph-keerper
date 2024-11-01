package app

import (
	"database/sql"
	_ "github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"log"
)

func Start(rootCmd *cobra.Command) {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func InitStorage() (*sql.DB, error) {
	// Открываем базу данных SQLite (если файла нет, он будет создан)
	db, err := sql.Open("sqlite", "./gophkeeper.db")
	if err != nil {
		return nil, err
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// SQL-запросы для создания таблиц
	schema := `
	CREATE TABLE IF NOT EXISTS 'app' (
	    id INTEGER UNIQUE DEFAULT 1,
		jwt TEXT NOT NULL,
		server_url TEXT
	);`

	// Выполнение SQL-запросов
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
