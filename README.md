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

## Commands

Available Commands:

`build`       Build the dev-local containers
`down`        Kill all running containsers
`exec`        Exec a command with the php container
`help`        Help about any command
`install`     Installs the base docker configuration into your application
`nginx-logs`  Tail nginx logs
`up`         A brief description of your command

## Connecting to MySQL
By default MySQL will be running on port `3306`. You can connect to the instance using the `root` user. The password is `password`. You will need to run `dev down` before running `dev up` on a different environment otherwise MySQL will clash (due to 3306 being taken by the old instance) and an error will be thrown.
