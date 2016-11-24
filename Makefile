TARGET := libmomentum_url_normalizer
all:	build

shared:
	go build -buildmode=c-shared -o $(TARGET).a ./main.go

test:
	go test ./lib/...

build:
	@$(MAKE) test
	@$(MAKE) shared

clean:
	rm -f $(TARGET).a $(TARGET).h
