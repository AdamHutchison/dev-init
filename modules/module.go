package modules

import (
	"github.com/spf13/cobra"
)

type Module struct {
	Identifier string
	Commands []*cobra.Command
}