package connect_to_server

import (
	"client-goph-keerper/internal/modules/connect_to_server/commands"
	"client-goph-keerper/internal/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var initAppCmd = &cobra.Command{
	Use:   "connect_to_server",
	Short: "Params for job application",
}

// StartsCmd инициализация команд - настроек клиента, для возможности начала работы.
func StartsCmd(s *storage.Storage) (*cobra.Command, error) {
	setServerCmd, err := commands.SetPathRemoteServerCommand(s)
	if err != nil {
		return nil, fmt.Errorf("error setting set server command: %w", err)
	}

	initAppCmd.AddCommand(setServerCmd)

	return initAppCmd, nil
}
