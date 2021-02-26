package base

import (
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "Launch docker containers defines in your ./docker-local/docker-compose.yml",
	Run: func(cmd *cobra.Command, args []string) {
		RmCmd.Run(cmd, args)
		command := "up -d"
		utils.DockerCompose(command)
	},
}

func init() {
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
