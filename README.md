# dev-init

Dev Init is a simple CLI tool that allows you to easily install and manage docker based dev environments for your PHP projects.

## Instalation

```bash
git clone git@github.com:AdamHutchison/dev-init.git
cd dev-init
make
```
Once you have run the Makefile check that the `dev` command is globally installed by running `dev`.

After you've confirmed `dev` is working, `cd` into a php project and run `dev install`. This will create a `docker-local` folder containing all the config you need for your dev environment.

## Base Commands

Available Commands:

Dev Init comes with some universal commands, these are based around managing docker containers and apply to all project types:

`build`       Build the dev-local containers

`down`        Kill all running containsers

`exec`        Exec a command with the php container

`help`        Help about any command

`install`     Installs the base docker configuration into your application

`up`         A brief description of your command

## Module

Using modules, Dev Init can be extended to provide useful project / framework specific commands. Modules are kept in the module folder and consist of a `module.go` file and then any specific cobra commands you want to be available for only specifc project types. `module.go` should define an instance of `modules.Module` from the `"github.com/AdamHutchison/dev-init/modules"` package. `module.Module` has two fields, firstly the `Identifier` which represents a file that identifies the project type. This file should always be present in the root folder of the project and should be unique to that project type e.g. Laravel projects always have an `artisan` file in the root directory which is used as an identifier to initiate the laravel commands. 
