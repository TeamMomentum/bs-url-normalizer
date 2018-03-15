// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"testing"
)

/* http/https プロトコルをhttpに統一するtest
* プロトコル正規化
https://example.com/ => http://example.com/

* パス末尾の正規化
http://example.com => http://example.com/
*/
func TestFirstNormalization(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		// TODO: Add test cases.

		{
			in:  "https://s.tabelog.com/aichi/A2303/A230302/23000859/?svd=20171120&svt=1900&svps=2&default_yoyaku_condition=1",
			out: "http://tabelog.com/aichi/A2303/A230302/23000859/",
		},
		{
			in:  "https://touch.pixiv.net/novel/show.php?id=8928225&uarea=tag",
			out: "http://www.pixiv.net/novel/show.php/?id=8928225&uarea=tag",
		},
		{
			in:  "https://touch.pixiv.net/bookmark.php?id=5020681&p=7",
			out: "http://www.pixiv.net/bookmark.php/",
		},
		{
			in:  "https://touch.pixiv.net/novel/recommend.php?id=6160013",
			out: "http://www.pixiv.net/novel/recommend.php/",
		},
		{
			in:  "https://example.com/",
			out: "http://example.com/",
		},
		{
			in:  "http://example.com",
			out: "http://example.com/",
		},
		{
			in:  "http://adm.shinobi.jp/a/66e78d8e9225eb41e4c240097cf56bb6?x=155&y=54&url=http%3A%2F%2Fgaingame.gendama.jp%2Fotenba%2Ftreasure&referrer=http%3A%2F%2Fgaingame.gendama.jp%2Fotenba%2Fgoal%2F2&user_id=&du=http%3A%2F%2Fgaingame.gendama.jp%2Fotenba%2Ftreasure&iw=1003&ih=995",
			out: "http://gaingame.gendama.jp/otenba/treasure/",
		},

		// New Line Tests
		//{
		//	in:  "http://example.com\n", => PANIC: invalid hostname
		//},
		{
			in:  "http://example.com/\n",
			out: "http://example.com/%0A/",
		},
		{
			in:  "http://example.com?id=hello\n",
			out: "http://example.com/?id=hello%0A",
		},
		{
			in:  "http://example.com?cb=hello\n",
			out: "http://example.com/",
		},
	}
	for _, tt := range tests {
		ul, err := url.Parse(tt.in)
		if err != nil {
			t.Errorf("%v : %v", tt.in, err)
		}
		t.Run(ul.Host+ul.Path, func(t *testing.T) {
			result := FirstNormalizeURL(ul)
			if result != tt.out {
				t.Errorf("unexpected %v : %v", result, tt.out)
			}
		})
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
