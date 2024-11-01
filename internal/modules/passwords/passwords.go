package passwords

import (
	"client-goph-keerper/internal/modules/passwords/commands"
	"github.com/spf13/cobra"
)

var pwdCmd = &cobra.Command{
	Use:   "passwords",
	Short: "Manage passwords",
}

func InitPwdCmd() *cobra.Command {
	savePwdCmd := commands.InitSavePwdCmd()
	deletePwdCmd := commands.InitDeletePwdCmd()
	getPwdCmd := commands.InitGetPwdCmd()
	allPwdCmd := commands.InitAllPwdCmd()
	updPwdCmd := commands.InitUpdatePwdCmd()

	pwdCmd.AddCommand(savePwdCmd)
	pwdCmd.AddCommand(deletePwdCmd)
	pwdCmd.AddCommand(getPwdCmd)
	pwdCmd.AddCommand(allPwdCmd)
	pwdCmd.AddCommand(updPwdCmd)

	return pwdCmd
}
