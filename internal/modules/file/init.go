package file

import "github.com/spf13/cobra"

func Init() *cobra.Command {
	// Флаги для команды file save
	saveFileCmd.Flags().String("path", "", "Path to the file")
	saveFileCmd.Flags().String("name", "", "Name of the file")
	saveFileCmd.MarkFlagRequired("path")
	saveFileCmd.MarkFlagRequired("name")

	// Флаги для команды file get
	getFileCmd.Flags().String("name", "", "Name of the file")
	getFileCmd.MarkFlagRequired("name")

	// Добавляем команды save и get к file
	fileCmd.AddCommand(saveFileCmd)
	fileCmd.AddCommand(getFileCmd)

	return fileCmd
}
