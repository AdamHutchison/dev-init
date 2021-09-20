package bootstrap

import (
	"github.com/AdamHutchison/dev-init/modules"
	r "github.com/AdamHutchison/dev-init/modules/register"
	"github.com/spf13/cobra"
)

func RegisterProjectCmds(cmd *cobra.Command) {
	module := r.GetContextModule()

	if module.Name != "" {
		registerCmds(r.GetCommonModule())
	}
	registerCmds(cmd, module)

}

func RegisterBaseCmds(cmd *cobra.Command) {
	registerCmds(cmd, r.GetBaseModule())
}

func registerCmds(cmd *cobra.Command, module modules.ModuleInterface) {
	for _, command := range module.GetCommands() {
		cmd.AddCommand(command)
	}
}
