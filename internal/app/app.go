package app

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
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

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
