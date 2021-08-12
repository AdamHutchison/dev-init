package zend

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Identifier: "config/modules.config.php",
	Commands: []*cobra.Command {
		InstallCmd,
		ReinstallCmd,
		php.TestCmd,
	},
}