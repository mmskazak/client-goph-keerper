package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the current user",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Здесь можно добавить логику для отправки запроса на выход
		req, err := http.NewRequest("GET", "http://localhost:8080/logout", nil)
		if err != nil {
			return err
		}

		// Добавьте заголовок авторизации, если требуется
		// req.Header.Set("Authorization", "Bearer "+token)

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

func InitLogoutCmd() *cobra.Command {
	// Можно добавить флаг для токена, если требуется
	return logoutCmd
}
