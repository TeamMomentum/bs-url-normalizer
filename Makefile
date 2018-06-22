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
	gometalinter -j 4 --deadline=300s ./... \
		--skip=main.go \
		--skip=/usr/local --skip=vendor \
		--exclude='/usr/local' --exclude='vendor/'
