package buffalo

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
		utils.Copy(utils.ResourceDir() + "/buffalo/docker-local", "./")

		goVersion, _ := cmd.Flags().GetString("go")

		utils.PrintTop("FROM golang:" + goVersion, utils.DockerFilePath("docker-local/conf/php/Dockerfile"))
		utils.PrintBottom("COMPOSE_PROJECT_NAME=" + utils.CurrentDirName(), ".env")
		utils.Link(utils.CurrentDir() + "/.env", utils.CurrentDir() + "/docker-local/")
	},
}

func init() {
	InstallCmd.Flags().String("go", "1.17.1", "The version of go to use")
	InstallCmd.MarkFlagRequired("go")
}