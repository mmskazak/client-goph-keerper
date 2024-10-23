package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var saveFileCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Saving file: %s with name: %s\n", path, name)
	},
}

// InitSaveCmdFile флаги для команды file save
func InitSaveCmdFile() *cobra.Command {
	saveFileCmd.Flags().String("path", "", "Path to the file")
	saveFileCmd.Flags().String("name", "", "Name of the file")
	saveFileCmd.MarkFlagRequired("path")
	saveFileCmd.MarkFlagRequired("name")
	return saveFileCmd
}
