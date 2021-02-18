package cmd

import (
	"fmt"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the dev-local containers",
	Run: func(cmd *cobra.Command, args []string) {
		buildCommand := []string{"--file", "./docker-local/docker-compose.yml", "build"}
		noCache, err := cmd.Flags().GetBool("no-cache")

		if err != nil {
			fmt.Println(utils.Fatal(err))
		}

		if noCache {
			buildCommand = append(buildCommand, "--no-cache")
		}

		utils.Exec("docker-compose", buildCommand...)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")
	buildCmd.Flags().BoolP("no-cache", "", false, "Build images from scratch")
}
