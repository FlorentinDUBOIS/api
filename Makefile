# Affect compiler and flags
CC := go build
GET := go get
GLIDE := glide
LINTER := gometalinter.v1
FORMAT := gofmt -w -s

GITHASH := $(shell git rev-parse HEAD)
CFLAGS := -ldflags "-X github.com/FlorentinDUBOIS/bouncer/cmd.githash=$(GITHASH)"

# build settings
BUILD_DIR := build
APP_NAME := bouncer
VPATH := $(BUILD_DIR)

# githash
GITHASH := $(shell git rev-parse HEAD)

.PHONY: build
build: bouncer.go
	$(CC) $(CFLAGS) -o $(BUILD_DIR)/$(APP_NAME) bouncer.go

.PHONY: get
get:
	$(GET) -u gopkg.in/alecthomas/gometalinter.v1 github.com/Masterminds/glide

.PHONY: install
install:
	$(GLIDE) install
	$(LINTER) --install --update

.PHONY: lint
lint:
	$(LINTER) --vendor --disable-all --enable=gocyclo --enable=structcheck --enable=aligncheck --enable=deadcode --enable=ineffassign --enable=dupl --enable=golint --enable=gotype --enable=goimports --enable=errcheck --enable=varcheck --enable=interfacer --enable=goconst --enable=gosimple --enable=staticcheck --enable=unparam --enable=unused --enable=misspell --enable=lll --enable=gas --tests ./...

.PHONY: format
format:
	$(FORMAT) ./cmd bouncer.go
