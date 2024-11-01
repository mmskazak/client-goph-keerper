package sync

import (
	"client-goph-keerper/internal/modules/sync/commands"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Update data from server",
}

func InitSyncCmd() *cobra.Command {
	syncPwdCmd := commands.InitSyncAllPwdCmd()
	syncCardsCmd := commands.InitSyncAllCardsCmd()
	syncAllFiles := commands.InitSyncAllFilesCmd()

	syncCmd.AddCommand(syncPwdCmd)
	syncCmd.AddCommand(syncCardsCmd)
	syncCmd.AddCommand(syncAllFiles)

	return syncCmd
}
