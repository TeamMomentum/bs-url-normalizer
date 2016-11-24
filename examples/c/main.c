#include <stdio.h>

#include "libmomentum_url_normalizer.h"

int main() {
	void *result;

	first_normalize_url("http://example.com/path/", &result);
	printf("%s\n", (char*)result);
	free_normalize_url(result);

	second_normalize_url("http://example.com/path/", &result);
	printf("%s\n", (char*)result);
	free_normalize_url(result);

	return 0;
}
