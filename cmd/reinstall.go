package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
	"os"
	"fmt"
)

// reinstallCmd represents the reinstall command
var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall the docker-local folder",
	Run: func(cmd *cobra.Command, args []string) {
		// Check data files exist
		if _, err := os.Stat(utils.DataDirPath()); os.IsNotExist(err) {
			fmt.Println(utils.Fatal("No data/ directory found at " + utils.DataDirPath() + " aborting...."))
			return
		}

		// backup data dir
		utils.Exec("cp", "-r", utils.DataDirPath(), utils.TmpDir())
		
		// remove original files
		utils.Exec("rm", "-rf", utils.CurrentDir() + "/docker-local")

		installCmd.Run(cmd, args)

		// replace data dir
		utils.Exec("cp", "-r", utils.TmpDir() + "/data", utils.CurrentDir() + "/docker-local")

		// remove  temp data backup
		utils.Exec("rm", "-rf", utils.TmpDir() + "/data")
	},
}

func init() {
	rootCmd.AddCommand(reinstallCmd)

	reinstallCmd.Flags().String("php", "8.0", "The version of php-fpm to use")
	reinstallCmd.MarkFlagRequired("php")
}
