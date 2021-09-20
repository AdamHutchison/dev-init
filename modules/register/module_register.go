package register

import (
	"os"

	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/context_modules/buffalo"
	"github.com/AdamHutchison/dev-init/modules/context_modules/laminas"
	"github.com/AdamHutchison/dev-init/modules/context_modules/laravel"
	"github.com/AdamHutchison/dev-init/modules/context_modules/wordpress"
	"github.com/AdamHutchison/dev-init/utils"
)

func GetContextModules() []modules.ModuleInterface {
	return []modules.ModuleInterface{
		laravel.Module,
		wordpress.Module,
		// zend.Module,
		laminas.Module,
		buffalo.Module,
	}
}

func GetContextModule() modules.ModuleInterface {
	for _, module := range GetContextModules() {

		identifier := utils.CurrentDir() + "/" + module.GetIdentifier()

		if _, err := os.Stat(identifier); err == nil {
			return module
		}
	}

	return modules.Module{}
}
