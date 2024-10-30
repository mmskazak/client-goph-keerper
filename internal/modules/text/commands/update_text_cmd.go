package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var updateTextCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a text entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		textID, _ := cmd.Flags().GetString("text_id")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		textContent, _ := cmd.Flags().GetString("text_content")

		data := map[string]string{
			"text_id":      textID,
			"title":        title,
			"description":  description,
			"text_content": textContent,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/text/update", bytes.NewBuffer(body))
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

func InitUpdateTextCmd() *cobra.Command {
	updateTextCmd.Flags().String("text_id", "", "Text ID to update")
	updateTextCmd.Flags().String("title", "", "New title for the text entry")
	updateTextCmd.Flags().String("description", "", "New description for the text entry")
	updateTextCmd.Flags().String("text_content", "", "New content of the text entry")
	updateTextCmd.MarkFlagRequired("text_id")
	updateTextCmd.MarkFlagRequired("title")
	updateTextCmd.MarkFlagRequired("description")
	updateTextCmd.MarkFlagRequired("text_content")
	return updateTextCmd
}
