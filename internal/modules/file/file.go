package file

import (
	"client-goph-keeper/internal/modules/file/commands"
	"github.com/spf13/cobra"
)

// fileCmd Команда для управления файлами (file)
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Manage files",
}

func InitFileCmd() *cobra.Command {
	saveFileCmd := commands.InitSaveCmdFile()
	getFileCmd := commands.InitGetFileCmd()
	fileCmd.AddCommand(saveFileCmd)
	fileCmd.AddCommand(getFileCmd)
	return fileCmd
}
