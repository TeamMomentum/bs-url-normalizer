package main

import (
	"net/url"

	"../../lib/urls"
)

func main() {
	u, _ := url.Parse("http://example.com/path/")
	println(urls.FirstNormalizeURL(u))
	println(urls.SecondNormalizeURL(u))
}
