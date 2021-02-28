package laravel

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/AdamHutchison/dev-init/installer"
)

// installCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the base docker configuration into your application",
	Run: func(cmd *cobra.Command, args []string) {
		installer.RunInstallCommand()
	},
}

func init() {
	InstallCmd.Flags().String("php", "8.0", "The version of php-fpm to use")
	InstallCmd.MarkFlagRequired("php")
}