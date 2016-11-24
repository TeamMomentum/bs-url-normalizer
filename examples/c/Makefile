all:	build

main:
	cp ../../libmomentum_url_normalizer.a ./
	cp ../../libmomentum_url_normalizer.h ./
	gcc -O2 -g -Wall -Wextra main.c -lmomentum_url_normalizer

build:
	cd ../../ && make
	@$(MAKE) main

clean:
	rm -rf a.out* libmomentum_url_normalizer.*
