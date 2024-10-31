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

func InitDB() (*sql.DB, error) {
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
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		jwt TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		description TEXT,
		path_to_file TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS cards (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		description TEXT,
		number INTEGER UNIQUE,
		pincode INTEGER,
		cvv INTEGER,
		expire DATE,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		login TEXT,
		password TEXT
	);

	CREATE TABLE IF NOT EXISTS texts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT,
		description TEXT,
		text TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	// Выполнение SQL-запросов
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
