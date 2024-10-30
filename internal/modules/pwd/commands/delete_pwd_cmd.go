package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deletePwdCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a password by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetInt("user_id")
		pwdID, _ := cmd.Flags().GetString("pwd_id")
		token, _ := cmd.Flags().GetString("token")

		data := map[string]interface{}{
			"user_id": userID,
			"pwd_id":  pwdID,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/delete", bytes.NewBuffer(body))
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

func InitDeletePwdCmd() *cobra.Command {
	deletePwdCmd.Flags().Int("user_id", 0, "User ID")
	deletePwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	deletePwdCmd.Flags().String("token", "", "Bearer token for authentication")
	deletePwdCmd.MarkFlagRequired("user_id")
	deletePwdCmd.MarkFlagRequired("pwd_id")
	deletePwdCmd.MarkFlagRequired("token")
	return deletePwdCmd
}
