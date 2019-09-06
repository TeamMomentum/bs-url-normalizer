all:	build

main:
	gcc -O2 -g -Wall -Wextra main.c -I$$PWD/../../build -L$$PWD/../../build -lmomentum_url_normalizer

build:
	$(MAKE) -C ../../ shared
	@$(MAKE) main

clean:
	rm -rf a.out*
