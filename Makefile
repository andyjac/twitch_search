BIN         = $(GOPATH)/bin
APP_NAME    = $(shell pwd | sed 's:.*/::')
APP_DIR     = $(GOPATH)/src/$(APP_NAME)
GLIDE      := $(shell command -v glide 2> /dev/null)
YARN       := $(shell command -v yarn 2> /dev/null)
NODE       := $(shell command -v node 2> /dev/null)

.PHONY: client
.PHONY: server
.PHONY: build
.PHONY: install

client:
	cd client && npm start

server:
	make build
	$(BIN)/$(APP_NAME)

build:
	go install

install:
ifdef	NODE
	@echo 'found node installation'
else
	@echo 'installing node...'
	curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.8/install.sh | bash
	nvm install node
endif

ifdef YARN
	@echo 'found yarn installation'
	cd client && yarn
else
	@echo 'installing yarn...'
	npm install -g yarn
	cd client && yarn
endif

ifdef GLIDE
	@echo 'found glide installation'
	glide install
else
	@echo 'installing glide...'
	curl https://glide.sh/get | sh
	glide install
endif
