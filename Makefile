BIN         = $(GOPATH)/bin
APP_NAME    = $(shell pwd | sed 's:.*/::')
APP_DIR     = $(GOPATH)/src/$(APP_NAME)
SETUP_DIR   = $(APP_DIR)/setup
GLIDE      := $(shell command -v glide 2> /dev/null)

.PHONY: client
.PHONY: server
.PHONY: setup
.PHONY: build
.PHONY: install

client:
	cd client && npm start

setup:
	make install

server:
	make build
	$(BIN)/$(APP_NAME)

build:
	go install

install:
	cd client && yarn install

ifdef GLIDE
	glide install
else
	$(warning "Skipping installation of Go dependencies: glide is not installed")
endif
