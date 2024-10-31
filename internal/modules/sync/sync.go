package sync

import (
	"client-goph-keerper/internal/modules/sync/commands"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Commands for syncing",
}

func InitSyncCmd() *cobra.Command {
	savePwdCmd := commands.InitSyncAllPwdCmd()

	syncCmd.AddCommand(savePwdCmd)

	return syncCmd
}
