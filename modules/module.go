package modules

import (
	"github.com/spf13/cobra"
)

type ModuleInterface interface {
	GetIdentifier() string
	GetCommands() []*cobra.Command
}

type Module struct {
	Identifier string
	Commands   []*cobra.Command
}

func (m Module) GetIdentifier() string {
	return m.Identifier
}

func (m Module) GetCommands() []*cobra.Command {
	return m.Commands
}
