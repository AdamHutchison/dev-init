package common

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/spf13/cobra"
)

var Module = modules.Module{
	Name: "base",
	Commands: []*cobra.Command{
		InstallCmd,
		ReinstallCmd,
		ExecCmd,
		SshCmd,
		UpCmd,
		BuildCmd,
		LogsCmd,
	},
}
