package zend

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// installCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the base docker configuration into your application",
	Run: func(cmd *cobra.Command, args []string) {
		// Copy docker files
		utils.Copy(utils.ResourceDir() + "/zend/docker-local", "./")

		php, _ := cmd.Flags().GetString("php")

		utils.PrintTop("FROM php:" + php + "-fpm", utils.PhpDockerFilePath())
		utils.PrintBottom("COMPOSE_PROJECT_NAME=" + utils.CurrentDirName(), ".env")
		utils.Link(utils.CurrentDir() + "/.env", utils.CurrentDir() + "/docker-local/")
	},
}

func init() {
	InstallCmd.Flags().String("php", "8.0", "The version of php-fpm to use")
	InstallCmd.MarkFlagRequired("php")
}