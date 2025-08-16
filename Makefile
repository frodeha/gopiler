.PHONY:= help
.DEFAULT_GOAL:= help

scratchpad: build ## Compile the scratchpad.go file
	$(build_target) ./debug/scratchpad

programs_dir=./test/programs
simple: build ## Compile the simple program
	$(build_target) $(programs_dir)/simple.go

lexer: build ## Compile the simple program
	$(build_target) lexer.go


programs: simple

build_target=./build/dasm
source_files=$(shell find . -iname "*.go")
build: $(source_files) ## Build the disassembler
	mkdir -p build
	go build -o build/dasm .

test: ## Run all tests
	go test ./...

help: ## Print this help menu
	@grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

