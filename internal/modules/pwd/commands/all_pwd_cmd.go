package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

var allPwdCmd = &cobra.Command{
	Use:   "all",
	Short: "List all passwords for a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetInt("user_id")
		token, _ := cmd.Flags().GetString("token")

		data := map[string]interface{}{
			"user_id": userID,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/all", bytes.NewBuffer(body))
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

		// Чтение тела ответа
		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %v", err)
		}

		fmt.Printf("Response: %s\n", responseData)
		return nil
	},
}

func InitAllPwdCmd() *cobra.Command {
	allPwdCmd.Flags().Int("user_id", 0, "User ID")
	allPwdCmd.Flags().String("token", "", "Bearer token for authentication")
	allPwdCmd.MarkFlagRequired("user_id")
	allPwdCmd.MarkFlagRequired("token")
	return allPwdCmd
}
