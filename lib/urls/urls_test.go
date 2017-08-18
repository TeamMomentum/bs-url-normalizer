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
