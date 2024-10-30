package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var getCardCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a card entry by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		cardID, _ := cmd.Flags().GetString("card_id")

		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/card/get/%s", cardID), nil)
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

func InitGetCardCmd() *cobra.Command {
	getCardCmd.Flags().String("card_id", "", "Card ID to retrieve")
	getCardCmd.MarkFlagRequired("card_id")
	return getCardCmd
}
