package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var updatePwdCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a password by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwdID, _ := cmd.Flags().GetString("pwd_id")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")

		data := map[string]interface{}{
			"pwd_id":      pwdID,
			"title":       title,
			"description": description,
			"credentials": map[string]string{
				"login":    login,
				"password": password,
			},
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/update", bytes.NewBuffer(body))
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

func InitUpdatePwdCmd() *cobra.Command {
	updatePwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	updatePwdCmd.Flags().String("title", "", "Title for the password entry")
	updatePwdCmd.Flags().String("description", "", "Description for the password entry")
	updatePwdCmd.Flags().String("login", "", "Login for the password entry")
	updatePwdCmd.Flags().String("password", "", "Password for the password entry")
	updatePwdCmd.MarkFlagRequired("pwd_id")
	updatePwdCmd.MarkFlagRequired("title")
	updatePwdCmd.MarkFlagRequired("login")
	updatePwdCmd.MarkFlagRequired("password")
	return updatePwdCmd
}
