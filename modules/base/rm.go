package base

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// rmCmd represents the rm command
var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove running containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers := utils.GetRunningContainers()
		downCmd.Run(cmd, args)
		arguments2 := append([]string{"rm"}, containers...)
		utils.Exec("docker", arguments2...)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
