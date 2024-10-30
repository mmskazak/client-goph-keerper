package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var getPwdCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwdID, _ := cmd.Flags().GetString("pwd_id")
		token, _ := cmd.Flags().GetString("token")

		url := fmt.Sprintf("http://localhost:8080/pwd/get/%s", pwdID)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		fmt.Printf("Response: %v\n", resp.Status)
		return nil
	},
}

func InitGetPwdCmd() *cobra.Command {
	getPwdCmd.Flags().String("pwd_id", "", "Password entry ID")
	getPwdCmd.Flags().String("token", "", "Bearer token for authentication")
	getPwdCmd.MarkFlagRequired("pwd_id")
	getPwdCmd.MarkFlagRequired("token")
	return getPwdCmd
}
