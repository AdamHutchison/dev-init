.DEFAULT_GOAL := all

INSTALL_DIRECTORY = ~/.config/dev-init
BIN_DIRECTORY = ~/.config/dev-init/bin/

all: copy_docker_files build_dev_init add_dev_init_to_path

copy_docker_files:
	@echo "Copying base docker files to $(INSTALL_DIRECTORY)"
	@if [ ! -d $(INSTALL_DIRECTORY) ]; then mkdir $(INSTALL_DIRECTORY); fi
	@cp -r ./resources/ $(INSTALL_DIRECTORY)

build_dev_init:
	@echo "Building dev tool"
	@go build -o dev
	@if [ ! -d $(BIN_DIRECTORY) ]; then mkdir $(BIN_DIRECTORY); fi
	@cp ./dev $(BIN_DIRECTORY)
	@rm ./dev

add_dev_init_to_path:
	@echo "Adding dev tool to path"
	@if [ $(type -P dev) ]; then echo "found"; fi
	@type -P "dev" && echo "dev command is already in PATH" || echo "export PATH=$(BIN_DIRECTORY):\$$PATH" >> ~/.zshrc