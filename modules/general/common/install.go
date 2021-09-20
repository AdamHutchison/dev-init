package common

import (
	"github.com/AdamHutchison/dev-init/modules/register"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the base docker configuration into your application",
	Run: func(cmd *cobra.Command, args []string) {
		module := register.GetContextModule()

		// Copy docker files
		utils.Copy(utils.ResourceDir() + module.GetName() + "docker-local", "./")

		version, _ := cmd.Flags().GetString("version")

		utils.PrintTop("FROM " + module.GetDockerImage() + version, utils.DockerFilePath(module.GetDockerFilePath()))
		utils.PrintBottom("COMPOSE_PROJECT_NAME=" + utils.CurrentDirName(), ".env")
		utils.Link(utils.CurrentDir() + "/.env", utils.CurrentDir() + "/docker-local/")
	},
}

func init() {
	InstallCmd.Flags().String("version", "", "Image version e.g. 8.0-fpm")
	InstallCmd.MarkFlagRequired("version")
}