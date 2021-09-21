package common

import (
	"strings"

	"github.com/AdamHutchison/dev-init/modules/register"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Exec a command with the php container",
	Run: func(cmd *cobra.Command, args []string) {
		module := register.GetContextModule()
		utils.DockerExec(module.GetDockerImage(), strings.Join(args, " "))
	},
}

func init() {
}
