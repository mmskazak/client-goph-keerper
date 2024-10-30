package card

import (
	"client-goph-keerper/internal/modules/card/commands"
	"github.com/spf13/cobra"
)

// Команда для управления паролями (pwd)
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Manage passwords",
}

// InitCardCommands инициализирует все команды управления картами
func InitCardCommands() *cobra.Command {
	cardCmd := &cobra.Command{
		Use:   "card",
		Short: "Card management commands",
	}

	saveCardCmd := commands.InitSaveCardCmd()
	deleteCardCmd := commands.InitDeleteCardCmd()
	getCardCmd := commands.InitGetCardCmd()
	updateCardCmd := commands.InitUpdateCardCmd()
	allCardCmd := commands.InitAllCardCmd()

	cardCmd.AddCommand(saveCardCmd)
	cardCmd.AddCommand(deleteCardCmd)
	cardCmd.AddCommand(getCardCmd)
	cardCmd.AddCommand(updateCardCmd)
	cardCmd.AddCommand(allCardCmd)

	return cardCmd
}
