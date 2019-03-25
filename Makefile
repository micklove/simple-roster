# Sample Makefile, for now
.DEFAULT_GOAL := run

ALL_PLATFORMS := linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x

PROJECT_NAME:=simple-roster
PROJECT_PKG_PATH:=github.com/micklove/$(PROJECT_NAME)
MAIN=cmd/$(PROJECT_NAME)/main.go

# TODO - use this pattern for multi os builds
example-build: 
	@echo "building GOOS:$(GOOS) GOARCH=$(GOARCH)" $*

build-%:
	@echo "task:  " $@
	@$(MAKE) example-build \
	  GOOS=$(firstword $(subst _, ,$*)) \
		GOARCH=$(lastword $(subst _, ,$*))

all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))

run: test
	@go run $(MAIN)

#run:
#	@go run cmd/simple-roster/main.go

install:
	@echo "install..."
	@go install -v -x ./...

build: clean
	@echo "build..."
#	@go build -v -x...
	@go build  ./...

clean: ## Remove all artifacts
	@printf "\n===================\nExecuting $@\n"
	@echo GOPATH: $(GOPATH)
	@go clean -x -i ./...
#	@rm -rf artifacts # clean CodeBuild artifacts

test: clean build  ## Run all unit tests
	@printf "\n===================\nExecuting $@\n"
	@go test ./...


