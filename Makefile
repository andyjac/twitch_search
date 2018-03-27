BIN         = $(GOPATH)/bin
APP_NAME    = $(shell pwd | sed 's:.*/::')
APP_DIR     = $(GOPATH)/src/$(APP_NAME)
SETUP_DIR   = $(APP_DIR)/setup
GLIDE      := $(shell command -v glide 2> /dev/null)

.PHONY: server
.PHONY: setup
.PHONY: client

client:
	cd client && npm start

setup:
	make install
	make migrate
	make seed

seed:
	go run $(SETUP_DIR)/seed.go

migrate:
	go run $(SETUP_DIR)/migrate.go

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
