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
