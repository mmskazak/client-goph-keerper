package commands

import (
	"client-goph-keerper/internal/storage"

	"fmt"
	"github.com/spf13/cobra"
)

// SetPathRemoteServerCommand NewSaveFileCmd Команда принимает базу данных в качестве параметра.
func SetPathRemoteServerCommand(s *storage.Storage) (*cobra.Command, error) {
	setServerCmd := &cobra.Command{
		Use:   "set",
		Short: "Set remote server url",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverURL, _ := cmd.Flags().GetString("server-url")

			if serverURL == "" {
				return fmt.Errorf("server-url flag is required")
			}

			// Сохраняем server_url в переданную базу данных
			_, err := s.DataBase.Exec(`
	   INSERT OR REPLACE
       INTO app_params (key, value) 
       VALUES (?, ?)`,
				"server_url", serverURL)
			if err != nil {
				return fmt.Errorf("failed to save server_url: %v", err)
			}

			fmt.Println("Server URL has been set successfully!")
			return nil

		},
	}
	setServerCmd.Flags().String("server-url", "", "Remote server url")
	err := setServerCmd.MarkFlagRequired("server-url")
	if err != nil {
		return nil, fmt.Errorf("failed to mark `server-url` flag as required")
	}
	return setServerCmd, nil
}
