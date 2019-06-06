RESOURCE_DIR=resources
ASSETS_DEST_DIR=lib
ASSETS_CMD=$(GOPATH)/bin/statik  # https://github.com/rakyll/statik
ASSETS_FILE=$(ASSETS_DEST_DIR)/assets/statik.go

TARGET := libmomentum_url_normalizer
all:	build

shared:
	go build -buildmode=c-shared -o $(TARGET).a ./main.go

assets:
	$(ASSETS_CMD) -p assets -src=$(RESOURCE_DIR) -dest=$(ASSETS_DEST_DIR)

$(ASSETS_FILE):
	@$(MAKE) assets

test: $(ASSETS_FILE)
	go test -v -race ./lib/...

build: $(ASSETS_FILE) test
	@$(MAKE) shared

clean:
	rm -f $(ASSETS_FILE)
	rm -f $(TARGET).a $(TARGET).h

lint:
	gometalinter -j 4 --deadline=300s ./... \
		--skip=main.go \
		--skip=/usr/local --skip=vendor \
		--exclude='/usr/local' --exclude='vendor/' \
		--cyclo-over=12
