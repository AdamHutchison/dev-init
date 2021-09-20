package buffalo

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Name: "buffalo",
	Identifier: ".buffalo.dev.yml",
	DockerImage: "golang",
	DockerFilePath: "docker-local/conf/go/Dockerfile",
	Commands: []*cobra.Command {
	},
}