package gophr

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/general/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Name: "gophr",
	Identifier: "config/application.config.php",
	DockerImage: "php",
	DockerFilePath: "docker-local/app/Dockerfile",
	Commands: []*cobra.Command {
		php.TestCmd,
	},
}