package pwd

import (
	"client-goph-keerper/internal/modules/pwd/commands"
	"github.com/spf13/cobra"
)

var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Manage passwords",
}

func InitPwdCmd() *cobra.Command {
	savePwdCmd := commands.InitSavePwdCmd()
	deletePwdCmd := commands.InitDeletePwdCmd()
	getPwdCmd := commands.InitGetPwdCmd()

	pwdCmd.AddCommand(savePwdCmd)
	pwdCmd.AddCommand(deletePwdCmd)
	pwdCmd.AddCommand(getPwdCmd)

	return pwdCmd
}
