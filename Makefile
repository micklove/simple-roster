# Sample Makefile, for now
.DEFAULT_GOAL := run

ALL_PLATFORMS := linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x

PROJECT_NAME:=simple-roster
PROJECT_PKG_PATH:=github.com/micklove/$(PROJECT_NAME)
MAIN=cmd/$(PROJECT_NAME)/main.go
COVERAGE_OUTPUT:=coverage.out
# TODO - use this pattern for multi os builds
example-build: 
	@echo "building GOOS:$(GOOS) GOARCH=$(GOARCH)" $*
ADDRESS:=localhost:8181
build-%:
	@echo "task:  " $@
	@$(MAKE) example-build \
	  GOOS=$(firstword $(subst _, ,$*)) \
		GOARCH=$(lastword $(subst _, ,$*))

all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))

run: test
	@go run $(MAIN) -addr=$(ADDRESS)

#run:
#	@go run cmd/simple-roster/main.go

install:
	@echo "install..."
	@go install -v -x ./...

build: clean
	@echo "build..."
#	@go build -v -x...
	@go build  ./...
	@go mod tidy
	@go mod verify

clean: ## Remove all artifacts
	@printf "\n===================\nExecuting $@\n"
	@echo GOPATH: $(GOPATH)
	@go clean -x -i ./...
#	@rm -rf artifacts # clean CodeBuild artifacts

test: clean build  ## Run all unit tests
	@printf "\n===================\nExecuting $@\n"
	@go test -v ./...

# See https://blog.golang.org/cover
cover: clean build ## Run coverage report and output to browser
	-rm $(COVERAGE_OUTPUT)
	@go test -v --coverprofile=$(COVERAGE_OUTPUT) ./...
	@go tool cover -html=$(COVERAGE_OUTPUT)

