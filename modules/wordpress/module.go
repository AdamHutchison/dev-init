package wordpress

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/modules"
)

var Module = modules.Module {
	Identifier: "wp-admin",
	Commands: []*cobra.Command {
		InstallCmd,
		ReinstallCmd,
		TestCmd,
	},
}