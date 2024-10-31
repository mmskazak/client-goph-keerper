package main

import (
	"client-goph-keerper/internal/app"
	"client-goph-keerper/internal/modules/auth"
	"client-goph-keerper/internal/modules/card"
	"client-goph-keerper/internal/modules/file"
	"client-goph-keerper/internal/modules/pwd"
	"client-goph-keerper/internal/modules/sync"
	"client-goph-keerper/internal/modules/text"
	"database/sql"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	// Инициализация базы данных
	db, err := app.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия соединения базы данныз: %v", err)
		}
	}(db)

	pwdCmd := pwd.InitPwdCmd()
	textCmd := text.InitTextCmd()
	cardCmd := card.InitCardCmd()
	fileCmd := file.InitFileCmd()
	syncCmd := sync.InitSyncCmd()
	authCmd := auth.InitAuthCmd()

	var rootCmd = &cobra.Command{Use: "app"}

	// Добавляем команды
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(textCmd)
	rootCmd.AddCommand(cardCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(authCmd)

	app.Start(rootCmd)
}
