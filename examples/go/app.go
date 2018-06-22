package main

import (
	"net/url"

	"github.com/TeamMomentum/bs-url-normalizer/lib/urls"
)

func main() {
	u, err := url.Parse("http://example.com/path/")
	if err != nil {
		panic(err)
	}
	println(urls.FirstNormalizeURL(u))
	println(urls.SecondNormalizeURL(u))
}
