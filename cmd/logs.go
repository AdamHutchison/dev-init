package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "nginx-logs",
	Short: "Tail nginx logs",
	Run: func(cmd *cobra.Command, args []string) {
		command := "docker-compose"
		arguments := []string{"--file", "docker-local/docker-compose.yml", "exec", "-T", "web", "tail", "-f", "/var/log/nginx/project_error.log"}

		utils.Exec(command, arguments...)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
