package urls

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	dokuhaPattern  = regexp.MustCompile(`/comicweb/viewer/comic/([^/]*)`)
	optimizeURLMap = map[string]func(*url.URL) bool{
		"dokuha.jp":         optimizeDokuhaURL,
		"novel.syosetu.org": createOptimizeURLFunc(regexp.MustCompile(`/\d+`)),
		"ncode.syosetu.com": createOptimizeURLFunc(regexp.MustCompile(`/[^/]+`)),
		"live.nicovideo.jp": optimizeLiveNicovideoURL,
		"s.maho.jp":         createOptimizeURLFunc(regexp.MustCompile(`/book/[^/]+/[^/]+/`)),
	}
)

func optimizeURL(ul *url.URL) {
	if cb, ok := optimizeURLMap[ul.Host]; ok {
		cb(ul)
	}
}

// 意味空間でURLを切り上げ、crawling対象のURLに変換する関数を返します
func createOptimizeURLFunc(re *regexp.Regexp) func(*url.URL) bool {
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
