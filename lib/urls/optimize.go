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
	optimizeURLMap = map[string]func(*url.URL) bool{
		"dokuha.jp":            optimizeDokuhaURL,
		"novel.syosetu.org":    createOptimizeURLCallBack(regexp.MustCompile(`/\d+`)),
		"ncode.syosetu.com":    createOptimizeURLCallBack(regexp.MustCompile(`/[^/]+`)),
		"live.nicovideo.jp":    optimizeLiveNicovideoURL,
		"s.maho.jp":            createOptimizeURLCallBack(regexp.MustCompile(`/book/[^/]+/[^/]+/`)),
		"enjoy.point.auone.jp": optimizeRestrictedURL("/gacha", "/reward", "/enquete"),
		"uranai.nosv.org":      optimizeRestrictedURL("/favorite.php"),
		"amigo.gesoten.com":    optimizeRestrictedURL(""),
		"gaingame.gesoten.com": optimizeRestrictedURL(""),
	}
)

func optimizeURL(ul *url.URL) {
	if cb, ok := optimizeURLMap[ul.Host]; ok {
		cb(ul)
	}
}

//Normalize 401 URLs
func optimizeRestrictedURL(prefixes ...string) func(*url.URL) bool {
	return func(ul *url.URL) bool {
		for _, prefix := range prefixes {
			if strings.HasPrefix(ul.Path, prefix) {
				ul.Path = prefix
				ul.RawQuery = ""
				return true
			}
		}
		return false
	}
}

// 意味空間でURLを切り上げ、crawling対象のURLに変換する関数を返します
func createOptimizeURLCallBack(re *regexp.Regexp) func(*url.URL) bool {
	return func(ul *url.URL) bool {
		groups := re.FindStringSubmatch(ul.Path)
		if len(groups) == 0 {
			return false
		}
		ul.Path = groups[0]
		return true
	}
}

/*
dokuha.jp用正規化関数
*/
func optimizeDokuhaURL(ul *url.URL) bool {
	groups := dokuhaPattern.FindStringSubmatch(ul.Path)
	if len(groups) == 0 {
		return false
	}
	ul.Path = "/comicweb/contents/comic/" + groups[1]
	return true
}

/*
	live.nicovideo.jp用正規化関数
*/
func optimizeLiveNicovideoURL(ul *url.URL) bool {
	if strings.HasPrefix(ul.Path, "/watch/") {
		ul.RawQuery = ""
		return true
	}
	return false
}
