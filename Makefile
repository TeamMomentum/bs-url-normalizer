RESOURCE_DIR=resources
ASSETS_DEST_DIR=lib
ASSETS_CMD=$(GOPATH)/bin/statik  # https://github.com/rakyll/statik

TARGET := libmomentum_url_normalizer
all:	build

shared:
	go build -buildmode=c-shared -o $(TARGET).a ./main.go

asset: $(ASSETS_FILE)
	$(ASSETS_CMD) -p assets -src=$(RESOURCE_DIR) -dest=$(ASSETS_DEST_DIR)

test:
	go test -v -race ./lib/...

build:
	@$(MAKE) assets
	@$(MAKE) test
	@$(MAKE) shared

clean:
	rm -f $(TARGET).a $(TARGET).h

lint:
	gometalinter -j 4 --deadline=300s ./... \
		--skip=main.go \
		--skip=/usr/local --skip=vendor \
		--exclude='/usr/local' --exclude='vendor/'
