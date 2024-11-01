package main

import (
	"client-goph-keerper/internal/app"
	"client-goph-keerper/internal/modules/auth"
	"client-goph-keerper/internal/modules/begin"
	"client-goph-keerper/internal/modules/file"
	"client-goph-keerper/internal/modules/passwords"
	"client-goph-keerper/internal/storage"
	"database/sql"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	// Инициализация SQLLite
	db, err := storage.Init()
	if err != nil {
		log.Fatalf("Init db err: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Close db err: %v", err)
		}
	}(db)

	installingCmd, err := begin.StartsCmd(db)
	if err != nil {
		log.Fatalf("install starts commands err: %v", err)
	}
	pwdCmd := passwords.InitPwdCmd()
	fileCmd := file.InitFileCmd()
	authCmd := auth.InitAuthCmd()

	var rootCmd = &cobra.Command{Use: "app"}

	// Добавляем команды
	rootCmd.AddCommand(installingCmd)
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(authCmd)

	app.Start(rootCmd)
}
