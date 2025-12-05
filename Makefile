SUBDIRS := $(filter-out days%,$(filter day%,$(wildcard */.)))
.PHONY: all plugins run clean $(SUBDIRS)

all: advent plugins
run: advent
	./advent

go.sum: go.mod
	go mod tidy

clean:
	rm -f advent plugins/*.so

advent: $(wildcard *.go utils/*.go days/*.go) go.sum
	go build -o advent .

plugins: ${SUBDIRS}
plugins/%: $(wildcard %/*.go) go.sum
	@echo "Building plugin for $(@D)"
	@go build -o $(@D)_plugin.so -buildmode=plugin $(patsubst plugins/%,./%,$(@D))

${SUBDIRS}: 
	@make plugins/$(dir $@:-=)
	