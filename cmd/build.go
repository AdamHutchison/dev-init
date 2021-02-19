package cmd

import (
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the dev-local containers",
	Run: func(cmd *cobra.Command, args []string) {
		buildCommand := "build"
		noCache, _ := cmd.Flags().GetBool("no-cache")

		if noCache {
			buildCommand += " --no-cache"
		}

		utils.DockerCompose(buildCommand)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")
	buildCmd.Flags().BoolP("no-cache", "", false, "Build images from scratch")
}
