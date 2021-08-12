package laravel

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Identifier: "artisan",
	Commands: []*cobra.Command {
		ArtCmd,
		InstallCmd,
		ReinstallCmd,
		RouteListCmd,
		php.TestCmd,
	},
}