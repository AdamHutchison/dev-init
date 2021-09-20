package bootstrap

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/general/base"
	"github.com/AdamHutchison/dev-init/modules/general/common"
	r "github.com/AdamHutchison/dev-init/modules/register"
	"github.com/spf13/cobra"
)

func RegisterProjectCmds(cmd *cobra.Command) {
	module := r.GetContextModule()

	if module.GetName() != "" {
		registerCmds(cmd, common.Module)
	}

	registerCmds(cmd, module)
}

func RegisterBaseCmds(cmd *cobra.Command) {
	registerCmds(cmd, base.Module)
}

func registerCmds(cmd *cobra.Command, module modules.ModuleInterface) {
	for _, command := range module.GetCommands() {
		cmd.AddCommand(command)
	}
}
