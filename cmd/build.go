package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the dev-local containers",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Exec("sudo", "chown", "-R", "adam:adam", utils.DataDirPath())
		utils.Exec("docker-compose", "--file", "./docker-local/docker-compose.yml", "build")
		utils.Exec("sudo", "chown", "-R", "adam:adam", utils.DataDirPath())
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
