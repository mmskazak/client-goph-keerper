package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deleteFileCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		fileID, _ := cmd.Flags().GetString("file_id")
		userID, _ := cmd.Flags().GetInt("user_id")
		token, _ := cmd.Flags().GetString("token")

		data := map[string]interface{}{
			"file_id": fileID,
			"user_id": userID,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/file/delete", bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
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

func InitDeleteFileCmd() *cobra.Command {
	deleteFileCmd.Flags().String("file_id", "", "File ID to delete")
	deleteFileCmd.Flags().Int("user_id", 0, "User ID")
	deleteFileCmd.Flags().String("token", "", "Bearer token for authentication")
	deleteFileCmd.MarkFlagRequired("file_id")
	deleteFileCmd.MarkFlagRequired("user_id")
	deleteFileCmd.MarkFlagRequired("token")
	return deleteFileCmd
}
