package passwords

import (
	"client-goph-keerper/internal/modules/passwords/commands"
	"client-goph-keerper/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var pwdCmd = &cobra.Command{
	Use:   "passwords",
	Short: "Manage passwords",
}

// InitPwdCmd инициализирует команду управления паролями
func InitPwdCmd(s *storage.Storage) (*cobra.Command, error) {
	// Инициализация команд
	savePwdCmd, err := commands.SetSavePasswordCmd(s)
	if err != nil {
		return nil, fmt.Errorf("ошибка настройки команды сохранения пароля: %v", err)
	}
	updPwdCmd, err := commands.SetUpdatePasswordCmd(s)
	if err != nil {
		return nil, fmt.Errorf("ошибка настройки команды обновления пароля: %v", err)
	}
	deletePwdCmd, err := commands.SetDeletePasswordCmd(s)
	if err != nil {
		return nil, fmt.Errorf("ошибка настройки команды удаления пароля: %v", err)
	}
	getPwdCmd, err := commands.SetGetPasswordCmd(s)
	if err != nil {
		return nil, fmt.Errorf("ошибка настройки команды получения пароля: %v", err)
	}
	allPwdCmd, err := commands.SetAllPasswordsCmd(s)
	if err != nil {
		return nil, fmt.Errorf("ошибка настройки команды получения всех паролей: %v", err)
	}

	// Добавление команд к главной команде
	pwdCmd.AddCommand(savePwdCmd)
	pwdCmd.AddCommand(updPwdCmd)
	pwdCmd.AddCommand(deletePwdCmd)
	pwdCmd.AddCommand(getPwdCmd)
	pwdCmd.AddCommand(allPwdCmd)

	return pwdCmd, nil
}
