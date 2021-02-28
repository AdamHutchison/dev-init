package installer

import (
	"github.com/AdamHutchison/dev-init/bootstrap"
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/utils"
	"os"
)

type InstallerInterface interface {
	GetModuleResourceFolder(m modules.ModuleInterface) string
}

type installer struct{}
 
func (i installer) RunInstallCommand(m modules.ModuleInterface) {
	for _, command := range m.GetCommands() {
		if command.Use == "install" {
			command.Run()
		}
	}
}

func (i installer) GetModuleResourceFolder(m modules.ModuleInterface) string {
	for _, module := range bootstrap.ModuleList {

		identifier := utils.CurrentDir() + "/" + module.GetIdentifier()

		if _, err := os.Stat(identifier); err == nil {
			return utils.ResourceDir() + "/" + m.GetName() + "/dev-local"
			break
		}
	}
}
