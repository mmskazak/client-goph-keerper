package app

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type GophKeeper struct {
	jwt string
	db  *sql.DB
}

func Start(pwdCmd *cobra.Command, fileCmd *cobra.Command) {
	var rootCmd = &cobra.Command{Use: "app"}

	// Добавляем команды pwd и file
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(fileCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func InitDB() (*sql.DB, error) {
	// Открываем базу данных SQLite (если файла нет, он будет создан)
	db, err := sql.Open("sqlite3", "./gophkeeper.db")
	if err != nil {
		return nil, err
	}
	log.Printf("SQLite DB created")

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Printf("SQLite DB opened")

	// SQL-запросы для создания таблиц
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
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
		user_id INTEGER,
		title TEXT,
		description TEXT,
		credentials JSON,
		FOREIGN KEY(user_id) REFERENCES users(id)
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
	log.Printf("Tables created successfully")

	return db, nil
}
