package php

import (
	"log"
	"github.com/AdamHutchison/dev-init/utils"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Exec a command with the php container",
	Run: func(cmd *cobra.Command, args []string) {
		command := "./vendor/bin/phpunit --colors"

		debug, _ := cmd.Flags().GetBool("debug")
		profile, _ := cmd.Flags().GetBool("profile")
		filter, _ := cmd.Flags().GetString("filter")
		stop, _ := cmd.Flags().GetBool("stop-on-failure")

		if (profile && debug) {
			log.Fatal(utils.Fatal("Both profile and debug flags passed to command. Only one allowed."))
		}

		addFilter(&command, filter)
		addDebug(&command, debug)
		addProfile(&command, profile)
		addStopOnFailure(&command, stop)

		utils.Exec("docker-compose", "--file", "docker-local/docker-compose.yml", "exec", "-T", "php", "/bin/bash", "-c", command)
	},
}

func init() {
	TestCmd.Flags().BoolP("debug", "d", false, "Enable step debugging using Xdebug")
	TestCmd.Flags().BoolP("profile", "p", false, "Generate profile using Xdebug")
	TestCmd.Flags().BoolP("stop-on-failure", "s", false, "Stop test run on first failure")
	TestCmd.Flags().String("filter", "", "Test class to filter by")
}

func addFilter(command *string, filter string) {
	if filter != "" {
		*command = *command + " --filter " + filter
	}
}

func addDebug(command *string, debug bool) {
	if debug {
		*command = "export XDEBUG_MODE=debug && " + *command
	}
}

func addProfile(command *string, debug bool) {
	if debug {
		*command = "export XDEBUG_MODE=profile && " + *command
	}
}

func addStopOnFailure(command *string, stop bool) {
	if stop {
		*command = *command + " --stop-on-failure"
	}
}