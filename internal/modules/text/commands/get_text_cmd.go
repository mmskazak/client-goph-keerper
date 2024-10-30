package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var getTextCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a text entry by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		textID, _ := cmd.Flags().GetString("text_id")

		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/text/get/%s", textID), nil)
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

func InitGetTextCmd() *cobra.Command {
	getTextCmd.Flags().String("text_id", "", "Text ID to retrieve")
	getTextCmd.MarkFlagRequired("text_id")
	return getTextCmd
}
