package bootstrap

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/laravel"
	"github.com/AdamHutchison/dev-init/modules/base"
)

var baseModule = base.Module

var moduleList = []modules.ModuleInterface{
	laravel.Module,
}