package laminas

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/modules"
)

var Module = modules.Module {
	Identifier: "config/modules.config.php",
	Commands: []*cobra.Command {
		InstallCmd,
		ReinstallCmd,
		TestCmd,
	},
}