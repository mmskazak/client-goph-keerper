package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deleteCardCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a card entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetInt("user_id")
		cardID, _ := cmd.Flags().GetString("card_id")

		data := map[string]interface{}{
			"user_id": userID,
			"card_id": cardID,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/card/delete", bytes.NewBuffer(body))
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

func InitDeleteCardCmd() *cobra.Command {
	deleteCardCmd.Flags().Int("user_id", 0, "User ID")
	deleteCardCmd.Flags().String("card_id", "", "Card ID to delete")
	deleteCardCmd.MarkFlagRequired("user_id")
	deleteCardCmd.MarkFlagRequired("card_id")
	return deleteCardCmd
}
