package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getPwdCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password",
	Run: func(cmd *cobra.Command, args []string) {
		login, _ := cmd.Flags().GetString("login")
		fmt.Printf("Getting password for login: %s\n", login)
	},
}

func InitGetPwdCmd() *cobra.Command {
	getPwdCmd.Flags().String("login", "", "Login for the password entry")
	getPwdCmd.MarkFlagRequired("login")
	return getPwdCmd
}
