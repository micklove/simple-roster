# Sample Makefile, for now

ALL_PLATFORMS := linux/amd64 linux/arm linux/arm64 linux/ppc64le linux/s390x

build: 
	@echo "building GOOS:$(GOOS) GOARCH=$(GOARCH)" $*

build-%:
	@echo "task:  " $@
	@$(MAKE) build \
	  GOOS=$(firstword $(subst _, ,$*)) \
		GOARCH=$(lastword $(subst _, ,$*))

all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))
