package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var allFilesCmd = &cobra.Command{
	Use:   "all",
	Short: "List all files for a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, _ := cmd.Flags().GetString("token")

		req, err := http.NewRequest("POST", "http://localhost:8080/file/all", nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		fmt.Printf("Response: %v\n", resp.Status)
		return nil
	},
}

func InitAllFilesCmd() *cobra.Command {
	allFilesCmd.Flags().String("token", "", "Bearer token for authentication")
	allFilesCmd.MarkFlagRequired("token")
	return allFilesCmd
}
