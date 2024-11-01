package begin

import (
	"client-goph-keerper/internal/modules/begin/commands"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
)

var initAppCmd = &cobra.Command{
	Use:   "begin",
	Short: "Params for job application",
}

func StartsCmd(db *sql.DB) (*cobra.Command, error) {
	setServerCmd, err := commands.SetServerCommand(db)
	if err != nil {
		return nil, fmt.Errorf("begin set server: %w", err)
	}

	initAppCmd.AddCommand(setServerCmd)

	return initAppCmd, nil
}
