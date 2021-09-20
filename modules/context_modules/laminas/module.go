package laminas

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/general/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Name: "laminas",
	Identifier: "config/modules.config.php",
	DockerImage: "php",
	DockerFilePath: "docker-local/conf/php/Dockerfile",
	Commands: []*cobra.Command {
		php.TestCmd,
	},
}