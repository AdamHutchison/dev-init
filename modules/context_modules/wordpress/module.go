package wordpress

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/general/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Name: "wordpress",
	Identifier: "wp-admin",
	DockerImage: "php",
	DockerFilePath: "docker-local/conf/php/Dockerfile",
	Commands: []*cobra.Command {
		ReinstallCmd,
		php.TestCmd,
	},
}