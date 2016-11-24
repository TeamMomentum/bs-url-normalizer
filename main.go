package main

import "C"

import (
	"net/url"

	"github.com/TeamMomentum/bscore/lib/utils/urls"
)

//export first_normalize_url
func first_normalize_url(raw string) string {
	ul, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	return urls.FirstNormalizeURL(ul)
}

//export second_normalize_url
func second_normalize_url(raw string) string {
	ul, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	return urls.SecondNormalizeURL(ul)
}

func main() {
}
