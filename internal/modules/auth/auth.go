package auth

import (
	"client-goph-keeper/internal/modules/auth/commands"
	"github.com/spf13/cobra"
)

// AuthCmd Команда для управления файлами (file)
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage auth",
}

func InitAuthCmd() *cobra.Command {
	registrationAuthCmd := commands.InitRegistrationCmdFlags()
	loginAuthCmd := commands.InitLoginCmdFlags()
	logoutAuthCmd := commands.InitLogoutCmdFlags()
	AuthCmd.AddCommand(registrationAuthCmd)
	AuthCmd.AddCommand(loginAuthCmd)
	AuthCmd.AddCommand(logoutAuthCmd)
	return AuthCmd
}
