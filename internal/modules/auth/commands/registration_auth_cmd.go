package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var registrationAuthCmd = &cobra.Command{
	Use:   "registration",
	Short: "Registration in app",
	Run: func(cmd *cobra.Command, args []string) {
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("Registering in app with login: %s and password: %s\n", login, password)
	},
}

// InitRegistrationCmdFlags добавление флагов для команды регистрации
func InitRegistrationCmdFlags() *cobra.Command {
	registrationAuthCmd.Flags().String("login", "", "Login for registration")
	registrationAuthCmd.Flags().String("password", "", "Password for registration")
	registrationAuthCmd.MarkFlagRequired("login")
	registrationAuthCmd.MarkFlagRequired("password")
	return registrationAuthCmd
}
