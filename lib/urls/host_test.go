// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

package urls

import (
	"testing"
)

//"net/url"
//"testing"

/*
パス階層レベルでの正規化のテスト
*/
/* func TestNormalizePathMap(t *testing.T) {
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
*/

func Test_normalizePath(t *testing.T) {
	tests := []struct {
		name     string
		rawurl   string
		isTarget bool
		want     string
	}{
		// ターゲットでないURLについては、後続処理で単純にホスト名を抽出する想定
		{"example.com (not normed)", "https://www.example.com/path/to/page.html?x=1&y=test", false, ""},

		// トピックは異なってもサイト全体で同一傾向と見なし、パス毎に正規化は行わない
		{"bakusai.com", "http://bakusai.com/areatop/acode=3/", false, ""},

		{"0bbs.jp", "http://0bbs.jp/2003v/", true, "http://0bbs.jp/2003v"}, // [DEPRECATED] 2017/01/31 サービス終了
		{"am-our.com", "https://am-our.com/idea/534/15188/", true, "https://am-our.com/idea"},
		{"ameblo.jp", "https://ameblo.jp/ebizo-ichikawa/entry-12387941308.html", true, "https://ameblo.jp/ebizo-ichikawa"},
		{"b.ibbs.info", "http://b.ibbs.info/sample/?anc=1-2#anc", true, "http://b.ibbs.info/sample"},
		{"bannch.com", "http://bannch.com/bs/bbs/798793/sub/index.html?test=123", true, "http://bannch.com/bs/bbs/798793"},
		{"bbs.mottoki.com", "http://bbs.mottoki.com/index?bbs=kinyuu&thread=&page=2", true, "http://bbs.mottoki.com/index?bbs=kinyuu"},
		{"bbs7.meiwasuisan.com", "http://bbs7.meiwasuisan.com/gravure/1529300473/", true, "http://bbs7.meiwasuisan.com/gravure"},
		// {"blg-girls.net", "blg-girls.net", true, ""},  // not found URL example, already closed?
		{"blog.dmm.co.jp", "http://blog.dmm.co.jp/actress/someone/", true, "http://blog.dmm.co.jp/actress"}, // 2017年サービス終了？
		{"blog.goo.ne.jp", "https://blog.goo.ne.jp/kuru0214/e/f0efbe813b94e456962c849b1f1c34f7", true, "https://blog.goo.ne.jp/kuru0214"},
		{"blog.kuruten.jp", "http://blog.kuruten.jp/katsutacompany77/411052", true, "http://blog.kuruten.jp/katsutacompany77"},
		{"blog.livedoor.jp", "http://blog.livedoor.jp/dqnplus/archives/1972090.html", true, "http://blog.livedoor.jp/dqnplus"},
		// {"blog.oricon.co.jp", "blog.oricon.co.jp", true, ""},　// not found URL example, already closed?
		{"blogs.yahoo.co.jp", "https://blogs.yahoo.co.jp/a209143707/66868842.html", true, "https://blogs.yahoo.co.jp/a209143707"},
		{"ch.nicovideo.jp", "http://ch.nicovideo.jp/horiemon/blomaga/201807", true, "http://ch.nicovideo.jp/horiemon"},
		{"cp.atrct.tv", "http://cp.atrct.tv/v/4zmEwAGgBg?pg=1", true, "http://cp.atrct.tv/v/4zmEwAGgBg"}, // アダルト画像掲示板？
		{"fanblogs.jp", "http://fanblogs.jp/30suppinnbihada/archive/512/0", true, "http://fanblogs.jp/30suppinnbihada"},
		{"free.jikkyo.org", "http://free.jikkyo.org/test/read.cgi/lsalofree/xxxxxxx", true, "http://free.jikkyo.org/test/read.cgi/lsalofree/xxxxxxx"}, // アダルトカテゴリ混在
		{"girlsnews.tv", "https://girlsnews.tv/unit/316023", true, "https://girlsnews.tv/unit"},                                                       // アダルトカテゴリ混在,
		// {"howcollect.jp", "http://howcollect.jp/tag/list/id/67", true, "http://howcollect.jp/tag/list/id/67"},                                         // TODO: 正規化階層要検討 (カテゴリ内の各記事は /article 直下, パス階層含む正規化不要？)
		// {"i.anisen.tv", "i.anisen.tv", true, ""}, // サイト閉鎖？
		{"ibbs.info", "http://ibbs.info/thread.php?no=1&id=TIEpachi", true, "http://ibbs.info/thread.php?id=TIEpachi"},
		{"jbbs.shitaraba.net", "https://jbbs.shitaraba.net/bbs/read.cgi/internet/0120/741222/", true, "https://jbbs.shitaraba.net/bbs/read.cgi/internet/0120/741222"}, // スレッドごとに正規化するのは細かすぎるかも？
		{"lineblog.me", "https://lineblog.me/kanosisters/archives/13190716.html", true, "https://lineblog.me/kanosisters"},
		{"lyze.jp", "https://lyze.jp/someone/diary/1/", true, "https://lyze.jp/someone"},
		//{"matome.naver.jp", "https://matome.naver.jp/odai/1234567890/", true, "https://matome.naver.jp/odai/1234567890"}, // TODO: 正規化階層要検討
		{"mbbs.tv", "http://mbbs.tv/u/?id=sample&p=2", true, "http://mbbs.tv/u?id=sample"},
		{"mblg.tv", "http://mblg.tv/someone/entry/26489/", true, "http://mblg.tv/someone"},
		{"mdpr.jp", "https://mdpr.jp/news/1234567890", true, "https://mdpr.jp/news"},
		//{"mess-y.com", "http://mess-y.com/", true, "http://mess-y.com"},	 // TODO: 正規化階層要検討 (/archive/*)
		// {"mikle.jp", "mikle.jp", true, ""},  // TODO: 正規化階層要検討
		{"nanos.jp", "http://nanos.jp/sample/novel/1/", true, "http://nanos.jp/sample"},
		//{"nikkan-spa.jp", "https://nikkan-spa.jp/1485749", true, "https://nikkan-spa.jp"}, // TODO: 正規化階層要検討 ( /${ARTICLE_ID} )
		{"plaza.rakuten.co.jp", "https://plaza.rakuten.co.jp/someone/diary/201807020000/", true, "https://plaza.rakuten.co.jp/someone"},
		{"rank.log2.jp", "http://rank.log2.jp/kro/category.php?cid=11", true, "http://rank.log2.jp/kro"}, //
		{"rara.jp", "https://rara.jp/someone/page231", true, "https://rara.jp/someone"},
		{"s1.ibbs.info", "http://s1.ibbs.info/thread.php?no=1&id=TIEpachi", true, "http://s1.ibbs.info/thread.php?id=TIEpachi"},
		{"s2.ibbs.info", "http://s2.ibbs.info/thread.php?no=1&id=TIEpachi", true, "http://s2.ibbs.info/thread.php?id=TIEpachi"},
		{"seesaawiki.jp", "http://seesaawiki.jp/foobar/d/FAQ", true, "http://seesaawiki.jp/foobar"},
		{"spora.jp", "http://spora.jp/mocchy/posts/606270", true, "http://spora.jp/mocchy"},
		// {"taishu.jp", "taishu.jp", true, ""}, // TODO: 正規化階層要検討 (/article/-/*)
		// {"talk.milkcafe.net", "talk.milkcafe.net", true, ""}, // milkcafe.net はそもそもサブドメインでカテゴリ分けしているので、パス正規化不要？
		// {"tocana.jp", "tocana.jp", true, ""}, // TODO: 正規化階層要検討
		{"woman.excite.co.jp", "https://woman.excite.co.jp/article/beauty/abc_def_00000/", true, "https://woman.excite.co.jp/article/beauty"},
		//{"www.asagei.com", "https://www.asagei.com/excerpt/106870", true, ""}, // TODO: 正規化階層要検討 ( /excerpt/* )
		// {"www.cyzo.com", "http://www.cyzo.com/2018/06/post_167645_entry.html", true, ""}, // TODO: 正規化階層要検討 ( /yyyy/mm/ )
		{"www.dclog.jp", "http://www.dclog.jp/someone/1/1234567890", true, "http://www.dclog.jp/someone"},
		{"www.ebbs.jp", "http://www.ebbs.jp/bbs.php?m=top&b=000000&guid=On", true, "http://www.ebbs.jp/bbs.php?b=000000"},
		{"www.eniblo.com", "https://www.eniblo.com/someone/2013/3/28/1364457329", true, "https://www.eniblo.com/someone"},
		//{"www.idolreport.jp", "www.idolreport.jp", true, ""}, // http://www.idolreport.jp/gravure/ しかない？
		//{"www.justhd.xyz", "www.justhd.xyz", true, ""}, // adult contents only?
		// {"www.menscyzo.com", "www.menscyzo.com", true, ""}, // // TODO: 正規化階層要検討 ( /yyyy/mm/ ), adult contents only?
		{"www.nikkan-gendai.com", "https://www.nikkan-gendai.com/articles/view/life/232439", true, "https://www.nikkan-gendai.com/articles/view/life"},
		{"www.tokyo-sports.co.jp", "https://www.tokyo-sports.co.jp/sports/othersports/000000/", true, "https://www.tokyo-sports.co.jp/sports"},
		{"www.zakzak.co.jp", "https://www.zakzak.co.jp/eco/news/180702/eco0000000-n1.html", true, "https://www.zakzak.co.jp/eco"},
		{"yaplog.jp", "http://yaplog.jp/someone/archive/0000", true, "http://yaplog.jp/someone"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul := mustURL(tt.rawurl)
			isNormalized := normalizePath(ul)
			if isNormalized != tt.isTarget {
				t.Errorf("normalizePath() returns %v, want %v", isNormalized, tt.isTarget)
			}
			if !isNormalized {
				return
			}
			if nURL := ul.String(); nURL != tt.want {
				t.Errorf("normalized URL = %v, want %v", nURL, tt.want)
			}
		})
	}
}
