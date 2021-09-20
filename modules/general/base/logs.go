package base

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// downCmd represents the down command
var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Follow the docker local compose logs",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DockerCompose("logs -f")
	},
}

func init() {

}
