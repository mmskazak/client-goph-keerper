package app

import (
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
