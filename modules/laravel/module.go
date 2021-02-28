package laravel

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/modules"
)

var Module = modules.Module {
	Name: "laravel",
	Identifier: "artisan",
	Commands: []*cobra.Command {
		ArtCmd,
		InstallCmd,
		ReinstallCmd,
		RouteListCmd,
		TestCmd,
	},
}