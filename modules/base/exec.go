package base

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"strings"
)

// execCmd represents the exec command
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Exec a command with the php container",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DockerExec(strings.Join(args, " "))
	},
}

func init() {
}
