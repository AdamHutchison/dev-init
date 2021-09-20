package utils

import (
	"fmt"
	"os/user"
	"os"
	"strings"
)

func HomeDir () string {
	usr, err := user.Current()

	if err != nil {
		fmt.Println(Fatal(err))
	}

	return usr.HomeDir
}

func CurrentDir() string  {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(Fatal(err))
	}

	return dir
}

func CurrentDirName() string  {
	split := strings.Split(CurrentDir(), "/")
	
	// return last item in slice as this is the name of the current folder
	return split[len(split)-1]
}

func ResourceDir() string {
	return HomeDir() + "/.config/dev-init/resources"
}

func TmpDir() string {
	return HomeDir() + "/.config/dev-init/tmp"
}

func DockerFilePath(relativePath string) string {
	return CurrentDir() + "/" + relativePath
}

func DataDirPath() string {
	return CurrentDir() + "/docker-local/data"
}

func DockerComposeFilePath() string {
	return CurrentDir() + "/docker-local/docker-compose.yml"
}