package urls

import (
	"net/url"
	"testing"
)

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
