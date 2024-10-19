package pwd

import "github.com/spf13/cobra"

// Команда для управления паролями (pwd)
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Manage passwords",
}
