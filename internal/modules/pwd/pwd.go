package pwd

import (
	"client-goph-keeper/internal/modules/pwd/commands"
	"github.com/spf13/cobra"
)

// Команда для управления паролями (pwd)
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Manage passwords",
}

func InitPwdCmd() *cobra.Command {
	savePwdCmd := commands.InitSavePwdCmd()
	getPwdCmd := commands.InitGetPwdCmd()
	pwdCmd.AddCommand(savePwdCmd)
	pwdCmd.AddCommand(getPwdCmd)
	return pwdCmd
}
