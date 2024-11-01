package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

var getPwdCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwdID, _ := cmd.Flags().GetString("pwd_id")

		url := fmt.Sprintf("http://localhost:8080/pwd/get/%s", pwdID)

		req, err := http.NewRequest("GET", url, nil)
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

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("Response: %v\n", resp.Status)
		fmt.Printf("Data: %v\n", fmt.Sprintf("%s", data))
		return nil
	},
}

func InitGetPwdCmd() *cobra.Command {
	getPwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	getPwdCmd.MarkFlagRequired("pwd_id")
	return getPwdCmd
}
