package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetRunningContainers() []string {
	out, err := exec.Command("docker", "ps", "-q").Output()

	if err != nil {
		fmt.Println(err)
	}

	split := strings.Split(string(out), "\n")

	if len(split) > 0 {
		split = split[:len(split)-1]
	}

	return split
}
