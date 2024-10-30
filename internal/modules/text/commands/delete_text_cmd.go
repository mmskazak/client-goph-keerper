package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deleteTextCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a text entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetInt("user_id")
		textID, _ := cmd.Flags().GetString("text_id")

		data := map[string]interface{}{
			"user_id": userID,
			"text_id": textID,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/text/delete", bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		// Замените на действительный токен
		req.Header.Set("Authorization", "Bearer YOUR_TOKEN_HERE")

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

func InitDeleteTextCmd() *cobra.Command {
	deleteTextCmd.Flags().Int("user_id", 0, "User ID")
	deleteTextCmd.Flags().String("text_id", "", "Text ID to delete")
	deleteTextCmd.MarkFlagRequired("user_id")
	deleteTextCmd.MarkFlagRequired("text_id")
	return deleteTextCmd
}
