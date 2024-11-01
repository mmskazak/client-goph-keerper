package main

import (
	"client-goph-keerper/internal/app"
	"client-goph-keerper/internal/config"
	"client-goph-keerper/internal/modules/auth"
	"client-goph-keerper/internal/modules/file"
	"client-goph-keerper/internal/modules/pwd"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	//Инициализация конфига
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	// Инициализация SQLLite
	db, err := app.InitStorage()
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
	fileCmd := file.InitFileCmd()
	authCmd := auth.InitAuthCmd()

	var rootCmd = &cobra.Command{Use: "app"}

	// Добавляем команды
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(authCmd)

	app.Start(rootCmd)
}
