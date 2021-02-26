package laravel

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
		utils.DockerExec("php artisan route:list | grep " + strings.Join(args, " "))
	},
}

func init() {
	// routeListCmd.PersistentFlags().String("foo", "", "A help for foo")
	// routeListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
