package cmd

import (
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Launch docker containers defines in your ./docker-local/docker-compose.yml",
	Run: func(cmd *cobra.Command, args []string) {
		rmCmd.Run(cmd, args)
		utils.Exec("docker-compose", "--file", "./docker-local/docker-compose.yml", "up", "-d", "--remove-orphans")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
