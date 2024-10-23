package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var logoutAuthCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Successfully logged out from the application.")
	},
}

// InitLogoutCmdFlags инициализация команды logout (без флагов)
func InitLogoutCmdFlags() *cobra.Command {
	// Никаких флагов не требуется для этой команды
	return logoutAuthCmd
}
