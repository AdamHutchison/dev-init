package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func Exec (cmd string, args ...string) {

	description := strings.Join(append([]string{"Running", cmd}, args...), " ")

	fmt.Println(Info(description))

	command := exec.Command(cmd, args...)

	stdout, err := command.StdoutPipe()

	command.Stderr = command.Stdout

	if err != nil {
		fmt.Println(Fatal(err))
		return
	}

	if err = command.Start(); err != nil {
		fmt.Println(Fatal(err))
		return
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	fmt.Println(Success("Command has completed successfully"))
}