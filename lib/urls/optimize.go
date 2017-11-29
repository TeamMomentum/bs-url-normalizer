// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	dokuhaPattern  = regexp.MustCompile(`/comicweb/viewer/comic/([^/]*)`)
	optimizeURLMap = map[string]func(*url.URL) *url.URL{
		"dokuha.jp":                      optimizeDokuhaURL,
		"novel.syosetu.org":              createOptimizeURLCallBack(regexp.MustCompile(`/\d+`)),
		"ncode.syosetu.com":              createOptimizeURLCallBack(regexp.MustCompile(`/[^/]+`)),
		"live.nicovideo.jp":              optimizeLiveNicovideoURL,
		"s.maho.jp":                      createOptimizeURLCallBack(regexp.MustCompile(`/book/[^/]+/[^/]+/`)),
		"enjoy.point.auone.jp":           optimizeRestrictedURL("/gacha", "/reward", "/enquete"),
		"uranai.nosv.org":                optimizeRestrictedURL("/favorite.php"),
		"amigo.gesoten.com":              optimizeRestrictedURL(""),
		"gaingame.gesoten.com":           optimizeRestrictedURL(""),
		"adm.shinobi.jp":                 parseAdframeURL("url"),
		"googleads.g.doubleclick.net":    parseAdframeURL("url"),
		"securepubads.g.doubleclick.net": parseAdframeURL("url"),
		"d.socdm.com":                    parseAdframeURL("tp"),
		"showads.pubmatic.com":           parseAdframeURL("pageURL"),
		"s.yimg.jp":                      parseAdframeURL("u"),
		"i.yimg.jp":                      parseAdframeURL("u"),
		"ssl.webtracker.jp":              parseAdframeURL("url"),
		"a.t.webtracker.jp":              parseAdframeURL("url"),
		"adw.addlv.smt.docomo.ne.jp":     parseAdframeURL("_url"),
	}
)

func optimizeURL(ul *url.URL) *url.URL {
	if cb, ok := optimizeURLMap[ul.Host]; ok {
		ul = cb(ul)
	}
	return ul
}

//Normalize Adframe URLs
func parseAdframeURL(key string) func(*url.URL) *url.URL {
	return func(ul *url.URL) *url.URL {
		if raw, ok := ul.Query()[key]; ok {
			u, err := url.Parse(raw[0])
			if err == nil {
				return u
			}
		}
		return ul
	}
}

//Normalize 401 URLs
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

// 意味空間でURLを切り上げ、crawling対象のURLに変換する関数を返します
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
dokuha.jp用正規化関数
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
	live.nicovideo.jp用正規化関数
*/
func optimizeLiveNicovideoURL(ul *url.URL) *url.URL {
	if strings.HasPrefix(ul.Path, "/watch/") {
		ul.RawQuery = ""
		return ul
	}
	return ul
}
