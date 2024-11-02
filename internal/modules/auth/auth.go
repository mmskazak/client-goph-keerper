package auth

import (
	"client-goph-keerper/internal/modules/auth/commands"
	"client-goph-keerper/internal/storage"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "User authentication commands",
}

// InitAuthCmd инициализирует команды аутентификации и добавляет их к authCmd
func InitAuthCmd(s *storage.Storage) (*cobra.Command, error) {
	// Инициализация команд с передачей объекта *storage.Storage
	registerCmd, err := commands.RegisterCommand(s)
	if err != nil {
		return nil, err
	}

	loginCmd, err := commands.LoginCommand(s)
	if err != nil {
		return nil, err
	}

	logoutCmd, err := commands.LogoutCommand(s)
	if err != nil {
		return nil, err
	}

	// Добавляем команды к authCmd
	authCmd.AddCommand(registerCmd)
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)

	return authCmd, nil
}
