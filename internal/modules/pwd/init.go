package pwd

import "github.com/spf13/cobra"

func Init() *cobra.Command {
	// Флаги для команды pwd save
	savePwdCmd.Flags().String("login", "", "Login for the password entry")
	savePwdCmd.Flags().String("password", "", "Password for the password entry")
	savePwdCmd.MarkFlagRequired("login")
	savePwdCmd.MarkFlagRequired("password")

	// Флаги для команды pwd get
	getPwdCmd.Flags().String("login", "", "Login for the password entry")
	getPwdCmd.MarkFlagRequired("login")

	// Добавляем команды save и get к pwd
	pwdCmd.AddCommand(savePwdCmd)
	pwdCmd.AddCommand(getPwdCmd)
	return pwdCmd
}
