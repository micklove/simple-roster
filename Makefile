# Sample Makefile, for now
.DEFAULT_GOAL := run

ALL_PLATFORMS := linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x

example-build: 
	@echo "building GOOS:$(GOOS) GOARCH=$(GOARCH)" $*

build-%:
	@echo "task:  " $@
	@$(MAKE) example-build \
	  GOOS=$(firstword $(subst _, ,$*)) \
		GOARCH=$(lastword $(subst _, ,$*))

all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))

run:
	@go run cmd/simple-roster/main.go

install:
	@go install ...

build:
	@go build ...

