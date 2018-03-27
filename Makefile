BIN      = $(GOPATH)/bin
APP_NAME = $(shell pwd | sed 's:.*/::')
APP_DIR  = $(GOPATH)/src/$(APP_NAME)
DB_DIR   = $(APP_DIR)/db
GLIDE    := $(shell command -v glide 2> /dev/null)

start:
	make build
	$(BIN)/$(APP_NAME)

build:
	go install

migrate:
	go run $(DB_DIR)/setup.go

install:
	@yarn install

ifdef GLIDE
	@glide install
else
	$(warning "Skipping installation of Go dependencies: glide is not installed")
endif
