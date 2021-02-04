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
	"strings"
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/utils"
)

// artCmd represents the art command
var artCmd = &cobra.Command{
	Use:   "art",
	Short: "Run php artisan command inside docker container",
	Run: func(cmd *cobra.Command, args []string) {
		command := "export XDEBUG_MODE=develop && php artisan " + strings.Join(args, " ")
		arguments := []string{"--file", "docker-local/docker-compose.yml", "exec", "-T", "php" ,"/bin/bash", "-c", command}
		utils.Exec("docker-compose", arguments...)

	},
}

func init() {
	rootCmd.AddCommand(artCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// artCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// artCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
