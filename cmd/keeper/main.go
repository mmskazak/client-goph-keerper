package main

import (
	"client-goph-keerper/internal/app"
	"client-goph-keerper/internal/modules/auth"
	"client-goph-keerper/internal/modules/begin"
	"client-goph-keerper/internal/modules/file"
	"client-goph-keerper/internal/modules/passwords"
	"client-goph-keerper/internal/storage"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	// Инициализация Storage
	s, err := storage.Init()
	if err != nil {
		log.Fatalf("Error init storage struct: %v", err)
	}

	installingCmd, err := begin.StartsCmd(s)
	if err != nil {
		log.Fatalf("Error install starts commands err: %v", err)
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
