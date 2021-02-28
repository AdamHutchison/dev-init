package modules

import (
	"github.com/spf13/cobra"
)

type ModuleInterface interface {
	GetName() string
	GetIdentifier() string
	GetCommands() []*cobra.Command
}

type Module struct {
	Name string
	Identifier string
	Commands   []*cobra.Command
}

func (m Module) GetName() string {
	return m.Name
}

func (m Module) GetIdentifier() string {
	return m.Identifier
}

func (m Module) GetCommands() []*cobra.Command {
	return m.Commands
}
