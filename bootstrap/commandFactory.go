package bootstrap

import (
	"github.com/AdamHutchison/dev-init/modules/base"
	"github.com/spf13/cobra"
)

func RegisterBaseCmds (cmd *cobra.Command) {
	for _, command := range base.Module.Commands {
		cmd.AddCommand(command)
	}
}