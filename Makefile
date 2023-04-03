RESOURCE_DIR=resources
ASSETS_DEST_DIR=lib
ASSETS_DIR=assets

TARGET := libmomentum_url_normalizer

export GO111MODULE := on

all:	build

shared:
	go build -buildmode=c-shared -o build/$(TARGET).a ./main.go

test: $(ASSETS_FILE)
	go test $(GOOPT) -v -race ./...

dep:
	go mod tidy

build: dep $(ASSETS_FILE) test
	@$(MAKE) shared

clean:
	rm -f $(ASSETS_FILE)
	rm -f $(TARGET).a $(TARGET).h

lint:
	golangci-lint run
