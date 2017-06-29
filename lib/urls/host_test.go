// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"testing"
)

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
