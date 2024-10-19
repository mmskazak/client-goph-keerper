package file

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

var getFileCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Getting file with name: %s\n", name)
	},
}
