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
		pwdID, _ := cmd.Flags().GetString("pwd_id")

		data := map[string]interface{}{
			"pwd_id": pwdID,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/delete", bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		// Получаем токен из базы данных
		token, err := getTokenFromDB()
		if err != nil {
			return fmt.Errorf("ошибка при получении токена: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)

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
	deletePwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	deletePwdCmd.MarkFlagRequired("pwd_id")
	return deletePwdCmd
}
