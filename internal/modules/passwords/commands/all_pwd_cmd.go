package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

var allPwdCmd = &cobra.Command{
	Use:   "all",
	Short: "List all passwords for a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		req, err := http.NewRequest("POST", "http://localhost:8080/pwd/all", nil)
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

		// Чтение тела ответа
		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ошибка чтения ответа: %v", err)
		}
		fmt.Printf("Status: %v\n", resp.Status)
		fmt.Printf("Response: %s\n", responseData)
		return nil
	},
}

func InitAllPwdCmd() *cobra.Command {
	return allPwdCmd
}
