package commands

import (
	"client-goph-keerper/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

// SetAllFilesCmd создает команду для получения списка всех файлов для пользователя
func SetAllFilesCmd(s *storage.Storage) (*cobra.Command, error) {
	allFilesCmd := &cobra.Command{
		Use:   "all",
		Short: "List all files for a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/file/all", s.ServerURL), nil)
			if err != nil {
				return fmt.Errorf("ошибка создания запроса: %v", err)
			}

			req.Header.Set("Authorization", s.Token)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("ошибка отправки запроса: %v", err)
			}
			defer resp.Body.Close()

			fmt.Printf("Response: %v\n", resp.Status)
			return nil
		},
	}

	// Установите токен как обязательный флаг
	err := allFilesCmd.MarkFlagRequired("token")
	if err != nil {
		return nil, fmt.Errorf("ошибка установки обязательного флага 'token': %v", err)
	}

	return allFilesCmd, nil
}
