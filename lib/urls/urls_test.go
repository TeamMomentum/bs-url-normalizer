package urls

import (
	"net/url"
	"testing"
)

// 不要なパラメータを除去できているかのtest
func TestRemoveQueryParameters(t *testing.T) {
	var (
		ul *url.URL
		nu string
	)

	ul = mustURL("http://blog.example.jp/tihoukoumu?utm_source=yahoo&utm_medium=cpc&utm_campaign=momentum&key=value")
	nu = FirstNormalizeURL(ul)
	stringCheck(t, "URL", "http://blog.example.jp/tihoukoumu/?key=value", nu)
}

/*
Query parameterの順序保証が保たれているかのtest
*元URL
http://example.com/tihoukoumu?d=1&a=2&c=3&b=4

*正規化後URL
http://example.com/tihoukoumu?a=2&b=4&c=3&d=1
*/
func TestQueryOrder(t *testing.T) {
	testURL := "http://example.com/tihoukoumu?d=1&a=2&c=3&b=4"
	results := make(map[string]bool)
	for i := 0; i < 100; i++ {
		nu := FirstNormalizeURL(mustURL(testURL))
		results[nu] = true
	}
	if len(results) != 1 {
		t.Error("URL query order should be stable.")
	}
	testURL = "http://example.com/tihoukoumu?a=2&c=3&b=4&d=1&utm_query=1"
	for i := 0; i < 100; i++ {
		nu := FirstNormalizeURL(mustURL(testURL))
		results[nu] = true
	}
	if len(results) != 1 {
		t.Error("URL query order should be stable.")
	}
}

/* http/https プロトコルをhttpに統一するtest
* プロトコル正規化
https://example.com/ => http://example.com/

* パス末尾の正規化
http://example.com => http://example.com/
*/
func TestFirstNormalization(t *testing.T) {
	var (
		ul *url.URL
		nu string
	)

	// プロトコル正規化
	ul = mustURL("https://example.com/")
	nu = FirstNormalizeURL(mustURL("https://example.com/"))

	if nu == ul.String() {
		t.Errorf("%v != %v", nu, ul.String())
	}

	// パス正規化
	ul = mustURL("http://example.com")
	nu = FirstNormalizeURL(mustURL("http://example.com"))

	if nu == ul.String() {
		t.Errorf("%v should not be %v", nu, ul.String())
	}
	if nu != "http://example.com/" {
		t.Errorf("%v should be %v", nu, "http://example.com/")
	}
}

/*
* パスの正規化
http://blog.livedoor.jp/tihoukoumu/sub/test => http://blog.livedoor.jp/tihoukoumu
*/
func TestSecondNormalization(t *testing.T) {
	ul := mustURL("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4")
	if sul := SecondNormalizeURL(ul); sul != "http://example.com" {
		t.Errorf("URL should not be %v", sul)
	}

	ul = mustURL("http://example.com:8000/tihoukoumu?d=1&a=2&c=3&b=4")
	if sul := SecondNormalizeURL(ul); sul != "http://example.com:8000" {
		t.Errorf("URL should not be %v", sul)
	}

	ul = mustURL("https://example.com:8000//tihoukoumu?d=1&a=2&c=3&b=4")
	if sul := SecondNormalizeURL(ul); sul != "http://example.com:8000" {
		t.Errorf("URL should not be %v", sul)
	}

	ul = mustURL("http://blog.livedoor.jp/tihoukoumu/sub/test?d=1&a=2&c=3&b=4")
	if sul := SecondNormalizeURL(ul); sul != "http://blog.livedoor.jp/tihoukoumu" {
		t.Errorf("URL should not be %v", sul)
	}

	ul = mustURL("http://bannch.com/a/b/c/d?d=1&a=2&c=3&b=4")
	if sul := SecondNormalizeURL(ul); sul != "http://bannch.com/a/b/c" {
		t.Errorf("URL should not be %v", sul)
	}
}

// パス階層分割関数のtest
func TestSplitNDomainPath(t *testing.T) {
	ul := mustURL("http://example.com/a/b/c/")
	if splitNDomainPath(ul, 2) != "example.com/a/b" {
		t.Errorf("%v != %v", splitNDomainPath(ul, 2), "example.com/a/b")
	}
}

/*
パス階層レベルでの正規化のテスト
*/
func TestNormalizePathMap(t *testing.T) {
	var (
		ul         *url.URL
		normalized bool
	)

	ul = mustURL("http://bannch.com/bs/bbs/798793/sub/index.html?test=123")
	normalized = normalizePath(ul)
	if !normalized {
		t.Errorf("%v should be normalized", ul.String())
		return
	}

	if ul.String() != "http://bannch.com/bs/bbs/798793" {
		t.Errorf("%v should be %v", ul.String(), "http://bannch.com/bs/bbs/798793")
		return
	}

	ul = mustURL("http://bbs.mottoki.com/index?bbs=kinyuu&thread=&page=2")
	normalized = normalizePath(ul)
	if !normalized {
		t.Errorf("%v should be normalized", ul.String())
		return
	}

	if ul.String() != "http://bbs.mottoki.com/index?bbs=kinyuu" {
		t.Errorf("%v should be %v", ul.String(), "http://bbs.mottoki.com/index?bbs=kinyuu")
		return
	}
}

func TestNormalizeMobileApp(t *testing.T) {
	var cases = []struct {
		rawurl string
		wants  string
	}{
		{
			rawurl: "https://itunes.apple.com/jp/app/minkara/id346528801?mt=8",
			wants:  "mobileapp::1-346528801",
		},
		{
			rawurl: "https://itunes.apple.com/app/id346528801",
			wants:  "mobileapp::1-346528801",
		},
		{
			rawurl: "https://play.google.com/store/apps/details?id=net.totopi.news&hl=jp",
			wants:  "mobileapp::2-net.totopi.news",
		},
		{
			rawurl: "https://play.google.com/store/apps/details?id=net.totopi.news",
			wants:  "mobileapp::2-net.totopi.news",
		},
	}

	for _, cs := range cases {
		ul := mustURL(cs.rawurl)
		if u := FirstNormalizeURL(ul); u != cs.wants {
			t.Errorf("%v != %v", u, cs.wants)
		}
	}
}

func TestOptimizeURL(t *testing.T) {
	var cases = []struct {
		rawurl string
		wants  string
	}{
		{
			"http://live.nicovideo.jp/watch/lv270002526?ref=notify&zroute=subscribe",
			"http://live.nicovideo.jp/watch/lv270002526",
		},
		{
			"http://ncode.syosetu.com/n7779dh/103/",
			"http://ncode.syosetu.com/n7779dh",
		},
		{
			"http://dokuha.jp/comicweb/viewer/comic/real/1",
			"http://dokuha.jp/comicweb/contents/comic/real",
		},
		{
			"https://novel.syosetu.org/81116/20.html",
			"https://novel.syosetu.org/81116",
		},
		{
			"http://s.maho.jp/book/2f7cc0g0b8fc434d/4767056014/2/",
			"http://s.maho.jp/book/2f7cc0g0b8fc434d/4767056014/",
		},
	}

	for _, cs := range cases {
		up, err := url.Parse(cs.rawurl)
		if err != nil {
			t.Error(err)
		}
		optimizeURL(up)
		if up.String() != cs.wants {
			t.Errorf("%v != %v", up.String(), cs.wants)
		}
	}
}

func stringCheck(t *testing.T, key, correct, other string) {
	if correct != other {
		t.Errorf("%v should be '%v', not '%v'.", key, correct, other)
	}
}

func mustURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}
