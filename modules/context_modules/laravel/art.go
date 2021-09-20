package laravel

import (
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
	"strings"
)

// artCmd represents the art command
var ArtCmd = &cobra.Command{
	Use:   "art",
	Short: "Run php artisan command inside docker container",
	Run: func(cmd *cobra.Command, args []string) {
		utils.DockerExec("php artisan " + strings.Join(args, " "))
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// artCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// artCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
