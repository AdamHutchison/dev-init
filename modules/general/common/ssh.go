package common

import (
	"os"
	"os/exec"

	"github.com/AdamHutchison/dev-init/modules/register"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Exec into the main project container",
	Run: func(cmd *cobra.Command, args []string) {
		service, _ := cmd.Flags().GetString("service")
		command := exec.Command("docker-compose", "--file", "docker-local/docker-compose.yml", "exec", service, "bash")

		command.Stdin = os.Stdin
		command.Stdout = os.Stdout

		err := command.Run()
		if err != nil {
			panic(err)
		}

		// reached once user types `exit` from the bash shell
		utils.Success("Docker: bash console ended successfully")
	},
}

func init() {
	module := register.GetContextModule()
	SshCmd.Flags().String("service", module.GetDockerImage(), "Docker compose service to ssh into")
}
