package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var getFileCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		fileID, _ := cmd.Flags().GetString("file_id")
		token, _ := cmd.Flags().GetString("token")

		req, err := http.NewRequest("GET", "http://localhost:8080/file/get/"+fileID, nil)
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

func InitGetFileCmd() *cobra.Command {
	getFileCmd.Flags().String("file_id", "", "File ID to retrieve")
	getFileCmd.Flags().String("token", "", "Bearer token for authentication")
	getFileCmd.MarkFlagRequired("file_id")
	getFileCmd.MarkFlagRequired("token")
	return getFileCmd
}
