package utils

import "strings"

func Copy(from string, to string) {
	Exec("cp", "-r", from, to)
}

func Delete(dir string) {
	Exec("rm", "-rf", dir)
}

func DockerExec(command string) {
	Exec("docker-compose", "--file", "docker-local/docker-compose.yml", "exec", "-T", "php", "/bin/bash", "-c", "export XDEBUG_MODE=develop && " + command)
}

func DockerCompose(command string) {
	baseCommand := []string{"--file", "docker-local/docker-compose.yml"}
	baseCommand = append(baseCommand, strings.Split(command, " ")...)
	Exec("docker-compose", baseCommand...)
}

func Link(from string, to string) {
	Exec("ln", "-s", from, to)
}

func PrintTop(text string, file string) {
	Exec("bash", "-c", "sed -i '1s/^/" + text + "'/ " + file)
}

func PrintBottom(text string, file string) {
	Exec("bash", "-c", "echo \"" + text + "\" >> " + file)
}
