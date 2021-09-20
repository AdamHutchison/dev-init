package zend

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"os"
	"fmt"
)

// reinstallCmd represents the reinstall command
var ReinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall the docker-local folder",
	Run: func(cmd *cobra.Command, args []string) {
		// Check data files exist
		if _, err := os.Stat(utils.DataDirPath()); os.IsNotExist(err) {
			fmt.Println(utils.Fatal("No data/ directory found at " + utils.DataDirPath() + " aborting...."))
			return
		}

		// backup data dir
		utils.Copy(utils.DataDirPath(), utils.TmpDir())
		
		// remove original files
		utils.Delete(utils.CurrentDir() + "/docker-local")

		InstallCmd.Run(cmd, args)

		// replace data dir
		utils.Copy(utils.TmpDir() + "/data", utils.CurrentDir() + "/docker-local")

		// remove  temp data backup
		utils.Delete(utils.TmpDir() + "/data")
	},
}

func init() {
	ReinstallCmd.Flags().String("php", "8.0", "The version of php-fpm to use")
	ReinstallCmd.MarkFlagRequired("php")
}
