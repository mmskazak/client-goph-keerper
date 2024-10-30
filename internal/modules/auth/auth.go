package auth

import (
	"client-goph-keerper/internal/modules/auth/commands"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "User authentication commands",
}

func InitAuthCmd() *cobra.Command {
	registerCmd := commands.InitRegisterCmd()
	loginCmd := commands.InitLoginCmd()
	logoutCmd := commands.InitLogoutCmd()

	authCmd.AddCommand(registerCmd)
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)

	return authCmd
}
