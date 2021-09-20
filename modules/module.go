package modules

import (
	"github.com/spf13/cobra"
)

type ModuleInterface interface {
	GetName() string
	GetIdentifier() string
	GetDockerFilePath() string
	GetDockerImage() string
	GetCommands() []*cobra.Command
}

type Module struct {
	Name string
	Identifier     string
	DockerImage    string
	DockerFilePath string
	Commands       []*cobra.Command
}

func (m Module) GetIdentifier() string {
	return m.Identifier
}

func (m Module) GetDockerFilePath() string {
	return m.DockerFilePath
}

func (m Module) GetDockerImage() string {
	return m.DockerImage
}

func (m Module) GetCommands() []*cobra.Command {
	return m.Commands
}
