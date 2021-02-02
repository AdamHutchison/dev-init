/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/AdamHutchison/dev-init/utils"
)

// execCmd represents the exec command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Exec a command with the php container",
	Run: func(cmd *cobra.Command, args []string) {

		command := "./vendor/bin/phpunit"

		arguments := []string{"--file", "docker-local/docker-compose.yml", "exec", "-T", "php", "/bin/bash", "-c"}

		debug, err := cmd.Flags().GetBool("debug")

		if err != nil {
			fmt.Println(utils.Fatal(err))
		}

		if debug != true {
			command = "export XDEBUG_MODE=develop && " + command
		}
		
		arguments = append(arguments, command)

		utils.Exec("docker-compose", arguments...)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().BoolP("debug", "d", false, "Enable Xdebug")

}
