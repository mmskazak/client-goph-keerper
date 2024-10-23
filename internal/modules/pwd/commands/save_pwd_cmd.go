package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var savePwdCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a password",
	Run: func(cmd *cobra.Command, args []string) {
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("Saving password for login: %s with password: %s\n", login, password)
	},
}

func InitSavePwdCmd() *cobra.Command {
	savePwdCmd.Flags().String("login", "", "Login for the password entry")
	savePwdCmd.Flags().String("password", "", "Password for the password entry")
	savePwdCmd.MarkFlagRequired("login")
	savePwdCmd.MarkFlagRequired("password")
	return savePwdCmd
}
