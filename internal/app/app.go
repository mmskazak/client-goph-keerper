package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type GophKeepet struct {
	jwt string
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
