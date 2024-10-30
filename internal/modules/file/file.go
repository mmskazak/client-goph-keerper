package file

import (
	"client-goph-keerper/internal/modules/file/commands"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Manage files",
}

func InitFileCmd() *cobra.Command {
	saveFileCmd := commands.InitSaveFileCmd()
	deleteFileCmd := commands.InitDeleteFileCmd()
	getFileCmd := commands.InitGetFileCmd()
	allFilesCmd := commands.InitAllFilesCmd()

	fileCmd.AddCommand(saveFileCmd)
	fileCmd.AddCommand(deleteFileCmd)
	fileCmd.AddCommand(getFileCmd)
	fileCmd.AddCommand(allFilesCmd)

	return fileCmd
}
