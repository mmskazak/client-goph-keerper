package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var loginAuthCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the application",
	Run: func(cmd *cobra.Command, args []string) {
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("Logging in with username: %s and password: %s\n", login, password)
	},
}

// InitLoginCmdFlags добавление флагов для команды входа
func InitLoginCmdFlags() *cobra.Command {
	loginAuthCmd.Flags().String("login", "", "Username for login")
	loginAuthCmd.Flags().String("password", "", "Password for login")
	loginAuthCmd.MarkFlagRequired("login")
	loginAuthCmd.MarkFlagRequired("password")
	return loginAuthCmd
}
