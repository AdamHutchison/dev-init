package bootstrap

import (
	"os"
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/base"
	"github.com/AdamHutchison/dev-init/modules/laravel"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

var moduleList = []modules.Module{
	laravel.Module,
}

func RegisterProjectCmds(cmd *cobra.Command) {
	for _, module := range moduleList {
		
		identifier := utils.CurrentDir() + "/" + module.Identifier

		if _, err := os.Stat(identifier); err == nil {
			registerCmds(cmd, module)
		}
	}
}

func RegisterBaseCmds(cmd *cobra.Command) {
	registerCmds(cmd, base.Module)
}

func registerCmds(cmd *cobra.Command, module modules.Module) {
	for _, command := range module.Commands {
		cmd.AddCommand(command)
	}
}
