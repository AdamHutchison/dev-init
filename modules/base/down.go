package base

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// downCmd represents the down command
var DownCmd = &cobra.Command{
	Use:   "down",
	Short: "Kill all running containsers",
	Run: func(cmd *cobra.Command, args []string) {
		containers := utils.GetRunningContainers()
		arguments := append([]string{"stop"}, containers...)
		utils.Exec("docker", arguments...)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
