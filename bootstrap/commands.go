package bootstrap

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
	"os"
)

func RegisterProjectCmds(cmd *cobra.Command) {
	for _, module := range ModuleList {

		identifier := utils.CurrentDir() + "/" + module.GetIdentifier()

		if _, err := os.Stat(identifier); err == nil {
			registerCmds(cmd, module)
			break
		}
	}
}

func RegisterBaseCmds(cmd *cobra.Command) {
	registerCmds(cmd, baseModule)
}

func registerCmds(cmd *cobra.Command, module modules.ModuleInterface) {
	for _, command := range module.GetCommands() {
		cmd.AddCommand(command)
	}
}
