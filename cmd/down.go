package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"os/exec"
	"strings"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Kill all running containsers",
	Run: func(cmd *cobra.Command, args []string) {
		out, err:= exec.Command("docker", "ps", "-q").Output()

		if err != nil {
			fmt.Println(err)
			return
		}

		split := strings.Split(string(out), "\n")

		if len(split) > 0 {
    		split = split[:len(split)-1]
		}

		arguments := append([]string{"stop"}, split...)

		fmt.Println("Stopping the following containers:", split)
		utils.Exec("docker", arguments...)
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
