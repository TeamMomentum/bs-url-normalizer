package main

import (
	"net/url"

	"github.com/TeamMomentum/bs-url-normalizer/lib/urls"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	funcs := map[string]interface{}{
		"URLParse":            url.Parse,
		"CrawlingURL":         urls.CrawlingURL,
		"FirstNormalizedURL":  urls.FirstNormalizeURL,
		"SecondNormalizedURL": urls.FirstNormalizeURL,
	}
	exports := js.Module.Get("exports")
	if exports != nil {
		exports.Set("BSURLNormalizer", funcs)
	} else {
		js.Global.Set("BSURLNormalizer", funcs)
	}
}
