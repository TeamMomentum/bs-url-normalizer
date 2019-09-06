RESOURCE_DIR=resources
ASSETS_DEST_DIR=lib
ASSETS_CMD=$(GOPATH)/bin/statik # https://github.com/rakyll/statik
ASSETS_DIR=assets
ASSETS_FILE=$(ASSETS_DEST_DIR)/$(ASSETS_DIR)/statik.go

TARGET := libmomentum_url_normalizer

HASH=$(shell git describe --tags)
GOENV=

all:	build

shared:
	go build -buildmode=c-shared -o $(TARGET).a ./main.go

assets:
	$(ASSETS_CMD) -f -p $(ASSETS_DIR) -src=$(RESOURCE_DIR) -dest=$(ASSETS_DEST_DIR)

$(ASSETS_FILE):
	@$(MAKE) assets

test: $(ASSETS_FILE)
	go test -v -race ./...

dep:
	dep ensure

build: dep $(ASSETS_FILE) test
	@$(MAKE) shared
	@$(MAKE) cmd

.PHONY: cmd
cmd: bs-url-normalizer
cmd_linux:
	$(eval GOENV=GOOS=linux GOARCH=amd64)
	$(MAKE) cmd GOENV="$(GOENV)"

bs-url-normalizer: cmd/bs-url-normalizer/main.go
	$(GOENV) go build -ldflags "-X 'main.version=$(HASH)'" ./cmd/bs-url-normalizer

install: cmd/bs-url-normalizer/main.go
	go install -ldflags "-X 'main.version=$(HASH)'" ./cmd/bs-url-normalizer

clean:
	rm -f $(ASSETS_FILE)
	rm -f $(TARGET).a $(TARGET).h

lint:
	golangci-lint run
