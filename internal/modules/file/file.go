package file

import "github.com/spf13/cobra"

// Команда для управления файлами (file)
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Manage files",
}
