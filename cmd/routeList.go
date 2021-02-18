package cmd

import (
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
	"strings"
)

// routeListCmd represents the routeList command
var routeListCmd = &cobra.Command{
	Use:   "route",
	Short: "List routes in Laravel project",
	Run: func(cmd *cobra.Command, args []string) {
		command := "export XDEBUG_MODE=develop && php artisan route:list | grep " + strings.Join(args, " ")
		arguments := []string{"--file", "docker-local/docker-compose.yml", "exec", "-T", "php", "/bin/bash", "-c", command}
		utils.Exec("docker-compose", arguments...)
	},
}

func init() {
	rootCmd.AddCommand(routeListCmd)
	// routeListCmd.PersistentFlags().String("foo", "", "A help for foo")
	// routeListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
