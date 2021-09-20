package buffalo

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Identifier: ".buffalo.dev.yml",
	Commands: []*cobra.Command {
		InstallCmd,
		ReinstallCmd,
	},
}