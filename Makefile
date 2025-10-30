.PHONY: all build clean vet test help test website
all: clean test build ## Run everything

BINDIR := ./bin
BINNAME := tryoutshell
BUILDFLAGS := -trimpath

clean: ## Clean the binary directory
	rm -rf $(BINDIR)

build: ## Build the binary
	CGO_ENABLED=0 go build $(BUILDFLAGS) -o $(BINDIR)/$(BINNAME) ./main.go

build-goreleaser: ## Build the binary using goreleaser
	goreleaser build --snapshot --clean

vet: ## Run go vet
	go vet ./...

test: ## Run go tests
	go test -v -coverprofile=profile.cov -covermode=atomic ./...

coverage: ## Show the coverage
	go tool cover -html=profile.cov

website:
	cd docs-website && yarn start

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
