package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

var (
	Info = Teal
	Warn = Yellow
	Fatal = Red
	Success = Green
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
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

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
		fmt.Sprint(args...))
	}

	return sprint
}