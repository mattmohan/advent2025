SUBDIRS := $(filter_out days,$(filter day%,$(wildcard */.)))
.PHONY: all plugins $(SUBDIRS)

all: advent plugins/day1.so

advent: $(wildcard *.go utils/*.go days/*.go)
	go build -o advent .

plugins/%.so: ${SUBDIRS}
	@echo "All plugins built for ${SUBDIRS}."

${SUBDIRS}: 
	@echo "Building plugin for $(@D)"
	@go build -o plugins/$(@D).so -buildmode=plugin ./$@