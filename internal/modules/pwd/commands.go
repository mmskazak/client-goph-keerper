package pwd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var savePwdCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a password",
	Run: func(cmd *cobra.Command, args []string) {
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("Saving password for login: %s with password: %s\n", login, password)
	},
}

var getPwdCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password",
	Run: func(cmd *cobra.Command, args []string) {
		login, _ := cmd.Flags().GetString("login")
		fmt.Printf("Getting password for login: %s\n", login)
	},
}
