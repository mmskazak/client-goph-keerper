package text

import (
	"client-goph-keerper/internal/modules/text/commands"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Manage text entries",
}

func InitTextCmd() *cobra.Command {
	saveTextCmd := commands.InitSaveTextCmd()
	deleteTextCmd := commands.InitDeleteTextCmd()
	getTextCmd := commands.InitGetTextCmd()
	updateTextCmd := commands.InitUpdateTextCmd()
	allTextCmd := commands.InitAllTextCmd()

	textCmd.AddCommand(saveTextCmd)
	textCmd.AddCommand(deleteTextCmd)
	textCmd.AddCommand(getTextCmd)
	textCmd.AddCommand(updateTextCmd)
	textCmd.AddCommand(allTextCmd)

	return textCmd
}
