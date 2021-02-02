package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"fmt"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Get a shell for the given container",
	Run: func(cmd *cobra.Command, args []string) {
		container, _ := cmd.Flags().GetString("container")
		fmt.Println(utils.Success("docker-compose --file docker-local/docker-compose.yml exec " + container + " bash"))
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	shellCmd.Flags().String("container", "php", "The container you need the shell for")
	shellCmd.MarkFlagRequired("php")
}
