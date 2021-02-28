package base

import (
	"github.com/spf13/cobra"
	"github.com/AdamHutchison/dev-init/modules"
)

var Module = modules.Module {
	Name: "base",
	Identifier: "",
	Commands: []*cobra.Command {
		BuildCmd,
		DownCmd,
		ExecCmd,
		RmCmd,
		UpCmd,
	},
}