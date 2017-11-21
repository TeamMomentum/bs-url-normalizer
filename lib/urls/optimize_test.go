// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
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
		{
			"https://enjoy.point.auone.jp/gacha/lottery/?token=+tnmAJSSfYGQ4k9ppHfZPHs2b5OWEYWOzbW1I4VkH",
			"https://enjoy.point.auone.jp/gacha",
		},
		{
			"https://enjoy.point.auone.jp/reward/?medid=walletmail&srcid=tameru&serial=0185&i=AeE9zC&ps=banner",
			"https://enjoy.point.auone.jp/reward",
		},
		{
			"https://enjoy.point.auone.jp/enquete/?aid=guronabi&bid=enquete&cid=",
			"https://enjoy.point.auone.jp/enquete",
		},
		{
			"http://uranai.nosv.org/recommend.php?urid=novel/flato",
			"http://uranai.nosv.org/recommend.php?urid=novel/flato",
		},
		{
			"http://uranai.nosv.org/favorite.php?crumb=0adf265b9d9c7921914c9f9fa32adeb3&add=novel/worldmadeh5&p=33&commu_id=worldmadehappye",
			"http://uranai.nosv.org/favorite.php",
		},
		{
			"http://amigo.gesoten.com/jewel/event/1494614261",
			"http://amigo.gesoten.com",
		},
		{
			"http://gaingame.gesoten.com/gaingame?user_id=000287102299&media_id=56&time=20171121001956&key=15E73E171612396096A3F68D8379841B",
			"http://gaingame.gesoten.com",
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
