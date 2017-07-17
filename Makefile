# Affect compiler
CC := go
GLIDE := glide
LINTER := gometalinter.v1
REMOVE := rm -r

# Command
BUILD := build
GET := get
FMT := gofmt

# Settings
BUILD_DIR := build
VPATH := $(BUILD_DIR)
APP_NAME := bouncer

# Shell
GITHASH := $(shell git rev-parse HEAD)

# Affect flags
CFLAGS := -ldflags "-X github.com/FlorentinDUBOIS/bouncer/cmd.githash=$(GITHASH)"

.PHONY: build
build: bouncer.go $(call rwildcard, ., *.go)
	$(CC) $(BUILD) $(CFLAGS) -o $(BUILD_DIR)/$(APP_NAME) bouncer.go

.PHONY: install
install:
	$(CC) $(GET) -u gopkg.in/alecthomas/gometalinter.v1 github.com/Masterminds/glide
	$(GLIDE) install
	$(LINTER) --install --update

.PHONY: format
format:
	$(FMT) -l -s -w ./cmd ./core bouncer.go

.PHONY: lint
lint:
	$(LINTER) --vendor --tests --deadline=180s --disable-all --enable=gofmt  ./...
	$(LINTER) --vendor --tests --deadline=180s --disable-all --enable=gas  ./...
	$(LINTER) --vendor --tests --deadline=180s --disable-all --enable=vet  ./...
	$(LINTER) --vendor --tests --deadline=180s --disable-all --enable=misspell  ./...
	$(LINTER) --vendor --tests --deadline=180s --disable-all --enable=golint  ./...
	$(LINTER) --vendor --tests --deadline=180s --disable-all --enable=ineffassign  ./...

.PHONY: clean
clean:
	$(REMOVE) -v -f $(BUILD_DIR)
