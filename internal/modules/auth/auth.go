package auth

import "github.com/spf13/cobra"

// Команда для управления файлами (file)
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage auth",
}
