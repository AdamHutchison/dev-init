package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Launch docker containers defines in your ./docker-local/docker-compose.yml",
	Run: func(cmd *cobra.Command, args []string) {
		command := "docker-compose"
		arguments := []string{"--file", "./docker-local/docker-compose.yml", "up", "-d" , "--remove-orphans"}

		utils.Exec(command, arguments...)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
