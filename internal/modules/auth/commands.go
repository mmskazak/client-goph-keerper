package auth

import (
	"fmt"
	"github.com/spf13/cobra"
)

var registrationAuthCmd = &cobra.Command{
	Use:   "registration",
	Short: "Registration",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Registration in app: %s with name: %s\n", path, name)
	},
}

var loginAuthCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Login to app: %s\n", name)
	},
}

var logoutAuthCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Logout from app: %s\n", name)
	},
}
