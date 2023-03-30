// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	itest5chDomain = "itest"
)

var (
	validate = validator.New()

	dokuhaPattern      = regexp.MustCompile(`/comicweb/viewer/comic/([^/]*)`)
	redirect5chPattern = regexp.MustCompile(`^/([^/]+)(/test/read.cgi/.*)$`)

	optimizeURLMap = map[string]func(*url.URL) *url.URL{
		"dokuha.jp":                       optimizeDokuhaURL,
		"live.nicovideo.jp":               optimizeLiveNicovideoURL,
		"s.maho.jp":                       createOptimizeURLCallBack(regexp.MustCompile(`/book/[^/]+/[^/]+/`)),
		"enjoy.point.auone.jp":            optimizeRestrictedURL("/gacha", "/reward", "/enquete"),
		"uranai.nosv.org":                 optimizeRestrictedURL("/favorite.php"),
		"amigo.gesoten.com":               optimizeRestrictedURL(""),
		"gaingame.gesoten.com":            optimizeRestrictedURL(""),
		"www.chatwork.com":                optimizeRestrictedURL(""),
		"adm.shinobi.jp":                  parseAdframeURL("url"),
		"googleads.g.doubleclick.net":     optimizeDoubleClickURL,
		"securepubads.g.doubleclick.net":  optimizeDoubleClickURL,
		"pubads.g.doubleclick.net":        optimizeDoubleClickURL,
		"d.socdm.com":                     optimizeSocdmURL,
		"showads.pubmatic.com":            parseAdframeURL("pageURL"),
		"s.yimg.jp":                       parseAdframeURL("u"),
		"i.yimg.jp":                       parseAdframeURL("u"),
		"ssl.webtracker.jp":               parseAdframeURL("url"),
		"a.t.webtracker.jp":               parseAdframeURL("url"),
		"adw.addlv.smt.docomo.ne.jp":      parseAdframeURL("_url"),
		"optimized-by.rubiconproject.com": parseAdframeURL("rf"),
		"jbbs.shitaraba.net":              parseAdframeURL("url"),
		"megalodon.jp":                    parseAdframeURL("url"),
		"ad.deqwas-dsp.net":               parseAdframeURL("url"),
		"krad20.deqwas.net":               parseAdframeURL("u"),
		"bidresult-dsp.ad-m.asia":         parseAdframeURL("rf"),
		"itest.5ch.net":                   optimizeItest5chURL,
		"itest.bbspink.com":               optimizeItest5chURL,
	}

	errEmptyURLString = errors.New("parse target URL is empty")
)

func optimizeURL(ul *url.URL) *url.URL {
	baseHost := ul.Host
	if cb, ok := optimizeURLMap[baseHost]; ok {
		ul = cb(ul)
		if ul.Host != baseHost {
			normalizeHost(ul)
		}
	}

	return ul
}

// Normalize Adframe URLs.
func parseAdframeURL(key string) func(*url.URL) *url.URL {
	return func(original *url.URL) *url.URL {
		if raw, ok := original.Query()[key]; ok {
			u, err := parsePotentialURL(raw[0])
			if err == nil {
				return u
			}
		}

		return original
	}
}

// Normalize 401 URLs.
func optimizeRestrictedURL(prefixes ...string) func(*url.URL) *url.URL {
	return func(ul *url.URL) *url.URL {
		for _, prefix := range prefixes {
			if strings.HasPrefix(ul.Path, prefix) {
				ul.Path = prefix
				ul.RawQuery = ""

				return ul
			}
		}

		return ul
	}
}

// 意味空間でURLを切り上げ、crawling対象のURLに変換する関数を返します.
func createOptimizeURLCallBack(re *regexp.Regexp) func(*url.URL) *url.URL {
	return func(ul *url.URL) *url.URL {
		groups := re.FindStringSubmatch(ul.Path)
		if len(groups) == 0 {
			return ul
		}

		ul.Path = groups[0]

		return ul
	}
}

/*
dokuha.jp用正規化関数.
*/
func optimizeDokuhaURL(ul *url.URL) *url.URL {
	groups := dokuhaPattern.FindStringSubmatch(ul.Path)
	if len(groups) == 0 {
		return ul
	}

	ul.Path = "/comicweb/contents/comic/" + groups[1]

	return ul
}

/*
live.nicovideo.jp用正規化関数.
*/
func optimizeLiveNicovideoURL(ul *url.URL) *url.URL {
	if strings.HasPrefix(ul.Path, "/watch/") {
		ul.RawQuery = ""

		return ul
	}

	return ul
}

/*
	d.socdm.com用正規化関数
*/
//nolint: cyclop
func optimizeSocdmURL(ul *url.URL) *url.URL {
	raw, ok := ul.Query()["sdktype"]
	if !ok {
		return ul
	}

	src := ""

	switch raw[0] {
	case "0":
		if raw, ok := ul.Query()["tp"]; ok {
			src = raw[0]
		}
	case "1":
		if raw, ok := ul.Query()["appbundle"]; ok {
			src = "mobileapp::2-" + raw[0]
		}
	case "2":
		if raw, ok := ul.Query()["appbundle"]; ok {
			src = "mobileapp::1-" + raw[0]
		}
	default:
		return ul
	}

	if len(src) == 0 {
		return ul
	}

	u, err := url.Parse(src)
	if err == nil {
		return u
	}

	return ul
}

/*
g.doubleclick.net用正規化関数.
*/
func optimizeDoubleClickURL(ul *url.URL) *url.URL {
	if raw, ok := ul.Query()["msid"]; ok {
		u, err := url.Parse("mobileapp::2-" + raw[0])
		if err == nil {
			return u
		}
	}

	if raw, ok := ul.Query()["_package_name"]; ok {
		u, err := url.Parse("mobileapp::1-" + raw[0])
		if err == nil {
			return u
		}
	}

	if raw, ok := ul.Query()["url"]; ok {
		u, err := url.Parse(raw[0])
		if err == nil {
			return u
		}
	}

	return ul
}

// optimize5chItestURL convert URL to where the page will be redirected by JavaScript
// e.g "http://itest.5ch.net/foo/test/read.cgi/abc/" => "http://foo.5ch.net/test/read.cgi/abc/"
func optimizeItest5chURL(ul *url.URL) *url.URL {
	if !strings.HasPrefix(ul.Host, itest5chDomain+".") { // URL is not target
		return ul
	}

	groups := redirect5chPattern.FindStringSubmatch(ul.Path)
	if len(groups) < 3 { //nolint: gomnd // unmatched pattern (unable to optimize)
		return ul
	}

	xul, err := url.Parse(ul.String()) // duplicate URL to preserve original one
	if err != nil {                    // unexpected error at this point
		return ul
	}

	xul.Host = strings.Replace(ul.Host, itest5chDomain, groups[1], 1)
	xul.Path = groups[2]

	if err := validate.Var(xul.String(), "url"); err != nil { // optimized URL is invalid
		return ul
	}

	return xul
}

//nolint:cyclop
func parsePotentialURL(rawurl string) (*url.URL, error) {
	if rawurl == "" {
		return nil, errEmptyURLString
	}

	parsed, parseErr := url.Parse(rawurl)
	if parseErr != nil { // assumed case: `hello.世界.com:8080/page.html` (multibyte hostname with port number)
		if host, tail, err := net.SplitHostPort(rawurl); err == nil && host != "" && tail != "" {
			return url.Parse("http://" + rawurl) //nolint: wrapcheck // retry with http scheme
		}

		return nil, fmt.Errorf("url.Parse: %w", parseErr)
	}

	// Scheme in the result of url.Parse would be lower case.
	// refer to go: https://github.com/golang/go/blob/9341fe073e6f7742c9d61982084874560dac2014/src/net/url/url.go#L528
	scheme := parsed.Scheme
	if scheme == "https" || scheme == "http" {
		return parsed, nil
	}

	if scheme == "" { // missing any scheme in rawurl
		return url.Parse("http://" + rawurl) //nolint: wrapcheck // re-parse with `http://` scheme prefix
	}

	if parsed.Host == "" { // case: scheme exists but missing authority part
		// Check if hostname was considered as the URL scheme
		// See also: https://github.com/golang/go/issues/12585
		if alt, err := url.Parse("http://" + rawurl); err == nil && parsed.Scheme == alt.Hostname() && alt.Port() != "" {
			return alt, nil
		}
	}

	return parsed, nil // non-http scheme URL (e.g. ftp://example.com/bar, data:,Hello%2C%20World!, etc)
}
