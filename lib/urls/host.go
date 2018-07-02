// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/lex/httplex"
)

var (
	// SPのHostとPCのHostの変換map
	spPCHostMap map[string]string
	// パス正規化対象ドメインと正規化関数のmap
	normalizePathMap map[string]func(*url.URL) bool
)

// normalizePath reduces known URLs to the top page of the website
func normalizePath(ul *url.URL) bool {
	f, ok := normalizePathMap[ul.Host]
	if !ok {
		return false
	}
	return f(ul)
}

func normalizePunycodeHost(ul *url.URL) {
	host, err := httplex.PunycodeHostPort(ul.Host)
	if err == nil {
		ul.Host = host
	}
}

// normalizeSPHost converts mobile URLs into their PC URLs.
func normalizeSPHost(ul *url.URL) {
	host, ok := spPCHostMap[ul.Host]
	if ok {
		ul.Host = host
	}
}

func makeStringStringMap(lines []string, sep string) map[string]string {
	m := make(map[string]string)
	for _, line := range lines {
		rows := strings.Split(line, sep)
		if len(rows) < 2 {
			continue
		}
		spHost := rows[0]
		pcHost := rows[1]
		m[spHost] = pcHost
	}
	return m
}

/* 指定したdomain => N階層のpairからパス正規化パターンを生成します
 */
func makeNormalizePathMap(lines []string, sep string) (map[string]func(*url.URL) bool, error) {
	m := make(map[string]func(*url.URL) bool)
	for _, line := range lines {
		// domain,パス階層,残したいparamIDの順に sep区切り
		rows := strings.Split(line, sep)
		if len(rows) < 3 {
			continue
		}
		domain := rows[0]
		num, err := strconv.Atoi(rows[1])
		if err != nil {
			return nil, err
		}
		param := rows[2]
		m[domain] = func(ul *url.URL) bool {
			ul.Path = splitNPath(ul, num)
			if param == "" {
				ul.RawQuery = ""
			} else {
				query := ul.Query()
				for key := range query {
					if param != key {
						query.Del(key)
					}
				}
				ul.RawQuery = query.Encode()
			}
			return true
		}
	}
	return m, nil
}

func init() {
	var err error
	spPCHostMap = makeStringStringMap(strings.Split(spHostData, "\n"), ",")
	normalizePathMap, err = makeNormalizePathMap(strings.Split(pathDepthData, "\n"), ",")
	if err != nil {
		panic(err)
	}
}

const spHostData = `sp.nicovideo.jp,www.nicovideo.jp
s.kakaku.com,kakaku.com
s.tabelog.com,tabelog.com
touch.pixiv.net,www.pixiv.net
touch.allabout.co.jp,allabout.co.jp
touch.navitime.co.jp,www.navitime.co.jp
a.excite.co.jp,www.excite.co.jp
sp.mainichi.jp,mainichi.jp
m.youtube.com,www.youtube.com
s.ameblo.jp,ameblo.jp
sp.okwave.jp,okwave.jp
sp.logsoku.com,www.logsoku.com
sp.daily.co.jp,www.daily.co.jp
sp.ultra-soccer.jp,web.ultra-soccer.jp
m.walkerplus.com,www.walkerplus.com
sp.bokete.jp,bokete.jp
sp.skincare-univ.com,www.skincare-univ.com
m.2log.sc,2log.sc
m.diodeo.jp,www.diodeo.jp
m.pideo.net,www.pideo.net
m.cinematoday.jp,www.cinematoday.jp
m.sponichi.co.jp,www.sponichi.co.jp`

const pathDepthData = `am-our.com,1,
ameblo.jp,1,
bannch.com,3,
blog.dmm.co.jp,1,
blog.goo.ne.jp,1,
blog.kuruten.jp,1,
blog.livedoor.jp,1,
blog.oricon.co.jp,1,
blogs.yahoo.co.jp,1,
ch.nicovideo.jp,1,
lineblog.me,1,
cp.atrct.tv,2,
fanblogs.jp,1,
free.jikkyo.org,4,
girlsnews.tv,2,
howcollect.jp,4,
i.anisen.tv,1,
jbbs.shitaraba.net,5,
lyze.jp,1,
mdpr.jp,1,
mess-y.com,4,
nanos.jp,1,
nikkan-spa.jp,1,
plaza.rakuten.co.jp,1,
seesaawiki.jp,1,
taishu.jp,2,
tocana.jp,1,
woman.excite.co.jp,2,
www.asagei.com,2,
www.cyzo.com,1,
www.eniblo.com,1,
www.idolreport.jp,1,
www.justhd.xyz,4,
www.menscyzo.com,1,
www.nikkan-gendai.com,3,
www.tokyo-sports.co.jp,2,
www.zakzak.co.jp,3,
yaplog.jp,1,
ibbs.info,1,id
s1.ibbs.info,1,id
s2.ibbs.info,1,id
b.ibbs.info,1,
talk.milkcafe.net,4,
mikle.jp,1,
matome.naver.jp,2,
bbs7.meiwasuisan.com,1,
mblg.tv,1,
blg-girls.net,2,
www.dclog.jp,1,
spora.jp,1,
0bbs.jp,1,
rank.log2.jp,1,
rara.jp,1,
bbs.mottoki.com,1,bbs
www.ebbs.jp,1,b
mbbs.tv,1,id
bakusai.com,5,`
