package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var updateCardCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a card entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		cardID, _ := cmd.Flags().GetString("card_id")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		number, _ := cmd.Flags().GetString("number")
		pincode, _ := cmd.Flags().GetString("pincode")
		cvv, _ := cmd.Flags().GetString("cvv")
		expire, _ := cmd.Flags().GetString("expire")

		data := map[string]string{
			"card_id":     cardID,
			"title":       title,
			"description": description,
			"number":      number,
			"pincode":     pincode,
			"cvv":         cvv,
			"expire":      expire,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/card/update", bytes.NewBuffer(body))
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

func InitUpdateCardCmd() *cobra.Command {
	updateCardCmd.Flags().String("card_id", "", "Card ID to update")
	updateCardCmd.Flags().String("title", "", "New title for the card entry")
	updateCardCmd.Flags().String("description", "", "New description for the card entry")
	updateCardCmd.Flags().String("number", "", "New card number")
	updateCardCmd.Flags().String("pincode", "", "New PIN code")
	updateCardCmd.Flags().String("cvv", "", "New CVV")
	updateCardCmd.Flags().String("expire", "", "New expiration date")
	updateCardCmd.MarkFlagRequired("card_id")
	updateCardCmd.MarkFlagRequired("title")
	updateCardCmd.MarkFlagRequired("description")
	updateCardCmd.MarkFlagRequired("number")
	updateCardCmd.MarkFlagRequired("pincode")
	updateCardCmd.MarkFlagRequired("cvv")
	updateCardCmd.MarkFlagRequired("expire")
	return updateCardCmd
}
