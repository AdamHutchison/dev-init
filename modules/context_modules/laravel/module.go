package laravel

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/general/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Name: "laravel",
	Identifier: "artisan",
	DockerImage: "php",
	DockerFilePath: "docker-local/conf/php/Dockerfile",
	Commands: []*cobra.Command {
		ArtCmd,
		ReinstallCmd,
		RouteListCmd,
		php.TestCmd,
	},
}