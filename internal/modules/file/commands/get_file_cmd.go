package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getFileCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Getting file with name: %s\n", name)
	},
}

func InitGetFileCmd() *cobra.Command {
	getFileCmd.Flags().String("name", "", "Name of the file")
	getFileCmd.MarkFlagRequired("name")

	return getFileCmd
}
