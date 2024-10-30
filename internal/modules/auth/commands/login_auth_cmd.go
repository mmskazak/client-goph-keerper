package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")

		data := map[string]string{
			"login":    login,
			"password": password,
		}

		body, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("ошибка кодирования JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")

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

func InitLoginCmd() *cobra.Command {
	loginCmd.Flags().String("login", "", "Login for the user")
	loginCmd.Flags().String("password", "", "Password for the user")
	loginCmd.MarkFlagRequired("login")
	loginCmd.MarkFlagRequired("password")
	return loginCmd
}
