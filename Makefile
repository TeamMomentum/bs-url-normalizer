RESOURCE_DIR=resources
ASSETS_DEST_DIR=lib
ASSETS_CMD=$(GOPATH)/bin/statik # https://github.com/rakyll/statik
ASSETS_DIR=assets
ASSETS_FILE=$(ASSETS_DEST_DIR)/$(ASSETS_DIR)/statik.go

TARGET := libmomentum_url_normalizer
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

clean:
	rm -f $(ASSETS_FILE)
	rm -f $(TARGET).a $(TARGET).h

lint:
	golangci-lint run
