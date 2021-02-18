package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"fmt"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the base docker configuration into your application",
	Run: func(cmd *cobra.Command, args []string) {
		// Copy docker files
		utils.Exec("cp", "-r", utils.ResourceDir() + "/docker-local", "./")

		php, err := cmd.Flags().GetString("php")

		if err != nil {
			fmt.Println(utils.Fatal(err))
		}

		utils.Exec("bash", "-c", "sed -i '1s/^/FROM php:" + php + "-fpm'/ " + utils.PhpDockerFilePath())

		// Add project name to env (used when naming containers)
		utils.Exec("bash", "-c", "echo \"COMPOSE_PROJECT_NAME=" + utils.CurrentDirName() + "\" >> .env")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	
	installCmd.Flags().String("php", "8.0", "The version of php-fpm to use")
	installCmd.MarkFlagRequired("php")
}