# dev-init

Dev Init is a simple CLI tool that allows you to easily install and manage docker based dev environments for several project types. Dev init will install a `docker-local` folder into your project that contains the full docker environment required to start developing. Dev-init also provides you with convenience commands to allow you to do common tasks from witin the containers.

## Dependencies
* Dev-init requires that go is installed on your machine. Installation details can be found [here](https://golang.org/doc/install).

## Definitions

* Context - this refers to the directory from where the `dev` command is being run from. `dev` will alter the commands available depending on which framework or project type it detects within the current directory.

* Module - A module refers to the driver code that allows dev init to be context aware (see the `Modules` section for more details).

## Instalation

```bash
git clone git@github.com:AdamHutchison/dev-init.git
cd dev-init
make
```
Once you have run the Makefile check that the `dev` command is globally installed by running `dev`.

After you've confirmed `dev` is working, `cd` into a php project and run `dev install`. This will create a `docker-local` folder containing all the config you need for your dev environment.

##  Commands

### Base Commands

These commands are available regardless of where you run the `dev` command from.

`down`        Kill all running containsers

`rm`          Remove all containers

### Common Commands

These commands are available from within any context:

`build`       Build the dev-local containers

`down`        Kill all running containsers

`exec`        Exec a command with the php container

`help`        Help about any command

`install`     Installs the base docker configuration into your application

`up`          Brings the docker files for the project up

`logs`        Shows the docker-compose logs for the project

`ssh`        SSH into a given docker compose container. Either provide the command with a `--service` (name of a service from `docker-local/docker-compose.yml`) flag or it'll default to the main docker compose service.

## Context Awareness

Dev-init is context aware. this means that the available `dev` commands will alter depending on which project type you are in. Currently the supported frameworks are:

* Laravel (PHP)
* Wordpress (PHP)
* Laminas (PHP)
* Buffalo (Golang)

To see the specific commands available for each framework checkout the `modules/context_modules` directory.

## Modules

### Overview
Dev-init achieves context awareness using modules. Modules can be thought of as drivers for different frameworks / project types. Dev Init can be extended to provide useful project / framework specific commands. A module, at it's core, consists of several cobra commands as well as a module specific instance of the `module.Module` struct.

The module defines all the specific information needed to register the module and for the module to run the base commands:

* Name - this is used to locate the `docker-local` folder when installing them into a project. The module specific `docker-local` folder must be placed in the `resources/{NAME}` directory. i.e. For the laravel module, the resources are located in the `resources/laravel/docker-local` folder.

* Identifier - this is a framework specific file or directory that dev-init will look for in order to determine what the project type is.

* Docker image - this defines the core docker image to be used in the docker file. A version must also be passed in on install. For laravel the docker image is `php` so running `dev install --version=8.0-fpm` would result in the `php:8.0-fpm` being used.

* Dockerfile path - dev-init needs to know the path of the main dockerfile in order to insert the correct base image to extend from at the top of it.

* A list of context specific commands. These will only be available in the projects where the `Identifier` exists.

```go
package laravel

import (
	"github.com/AdamHutchison/dev-init/modules"
	"github.com/AdamHutchison/dev-init/modules/general/php"
	"github.com/spf13/cobra"
)

var Module = modules.Module {
	Name: "laravel",
	Identifier: "artisan",
	DockerImage: "php",
	DockerFilePath: "docker-local/conf/php/Dockerfile",
	Commands: []*cobra.Command {
		ArtCmd,
		RouteListCmd,
		php.TestCmd,
	},
}
```

### The Base Module (Digging Deeper)

The base module defines commands that are always available.

### The Common Module
The common module defines commands that are availavble from within any context (but only if thwre is a context!).

### Module Registration
Once you have created a module it must be registerd within the `GetContextModules()` function of the `github.com/AdamHutchison/dev-init/modules/register` package.
