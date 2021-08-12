package laminas

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// execCmd represents the exec command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Exec a command with the php container",
	Run: func(cmd *cobra.Command, args []string) {
		command := "./vendor/bin/phpunit"

		debug, _ := cmd.Flags().GetBool("debug")
		filter, _ := cmd.Flags().GetString("filter")
		stop, _ := cmd.Flags().GetBool("stop-on-failure")

		addFilter(&command, filter)
		addDebug(&command, debug)
		addStopOnFailure(&command, stop)

		utils.Exec("docker-compose", "--file", "docker-local/docker-compose.yml", "exec", "-T", "php", "/bin/bash", "-c", command)
	},
}

func init() {
	TestCmd.Flags().BoolP("debug", "d", false, "Enable Xdebug")
	TestCmd.Flags().BoolP("stop-on-failure", "s", false, "Stop test run on first failure")
	TestCmd.Flags().String("filter", "", "Test class to filter by")
}

func addFilter(command *string, filter string) {
	if filter != "" {
		*command = *command + " --filter " + filter
	}
}

func addDebug(command *string, debug bool) {
	if !debug {
		*command = "export XDEBUG_MODE=develop && " + *command
	}
}

func addStopOnFailure(command *string, stop bool) {
	if stop {
		*command = *command + " --stop-on-failure"
	}
}