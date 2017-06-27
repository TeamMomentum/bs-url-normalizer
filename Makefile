TARGET := libmomentum_url_normalizer
all:	build

shared:
	go build -buildmode=c-shared -o $(TARGET).a ./main.go

test:
	go test -v -race ./lib/...

build:
	@$(MAKE) test
	@$(MAKE) shared

clean:
	rm -f $(TARGET).a $(TARGET).h

lint:
	golint -set_exit_status $$(go list ./... | grep -v vendor)
	(! gofmt -s -d $$(find ./ -name "*.go" | grep -v vendor | grep -v assets.go) | grep '^')
	unused $$(go list ./... | grep -v vendor)
	gosimple $$(go list ./... | grep -v vendor)
	errcheck $$(go list ./... | grep -v vendor)
	staticcheck $$(go list ./... | grep -v vendor)
	go vet $$(go list ./... | grep -v /vendor/)
