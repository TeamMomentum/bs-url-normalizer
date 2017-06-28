// Copyright 2016 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultScheme = "http"
)

var (
	disusedParametersMap map[string]bool
	// SPのHostとPCのHostの変換map
	spPCHostMap map[string]string
	// パス正規化対象ドメインと正規化関数のmap
	normalizePathMap map[string]func(*url.URL) bool
)

type Normalizer struct {
	cURL  string
	n1URL string
	n2URL string
	url   *url.URL
}

func NewNormalizer(ul *url.URL) (n *Normalizer, err error) {
	if !isValidScheme(ul) {
		n = nil
		err = fmt.Errorf("invalid scheme %v", ul.Scheme)
		return
	}

	url1 := ""
	url2 := ""
	if isStaticURL(ul) {
		url1 = ul.String()
		url2 = url1
	}

	n = &Normalizer{
		url:   ul,
		n1URL: url1,
		n2URL: url2,
	}

	return
}

func (n *Normalizer) CrawlingURL() string {
	if n.cURL == "" {
		optimizeURL(n.url)
		n.cURL = n.url.String()
	}
	return n.cURL
}

func (n *Normalizer) FirstNormalizedURL() string {
	if n.n1URL != "" {
		return n.n1URL
	}

	ul := n.url
	n.CrawlingURL()

	n.n1URL = normalizeMobileAppURL(ul)
	if n.n1URL != "" {
		return n.n1URL
	}

	removeQueryParameters(ul, ul.Query())
	normalizeSPHost(ul)
	normalizeScheme(ul)
	normalizePathSuffix(ul)
	n.n1URL = ul.String()

	return n.n1URL
}

func (n *Normalizer) SecondNormalizedURL() string {
	n.FirstNormalizedURL()
	if n.n2URL != "" {
		return n.n2URL
	}

	if normalizePath(n.url) {
		n.n2URL = n.url.String()
	} else {
		n.n2URL = n.url.Scheme + "://" + n.url.Host
	}
	return n.n2URL
}

// removeQueryParameters removes the unnecessary query parameters.
func removeQueryParameters(ul *url.URL, query url.Values) {
	deleteQueries := []string{}
	for key := range query {
		if strings.HasPrefix(key, "utm_") {
			deleteQueries = append(deleteQueries, key)
		} else if _, ok := disusedParametersMap[key]; ok {
			deleteQueries = append(deleteQueries, key)
		}
	}

	if len(deleteQueries) != 0 {
		for _, key := range deleteQueries {
			query.Del(key)
		}
	}

	ul.RawQuery = query.Encode()
}

// IsValidURL
func IsValidURL(ul *url.URL) bool {
	for _, s := range supportedSchemes {
		if ul.Scheme == s {
			return true
		}
	}
	return false
}

// normalizeScheme replaces all schemes into http
func normalizeScheme(ul *url.URL) {
	ul.Scheme = defaultScheme
}

// normalizePathSuffix keeps the end of URL as '/'.
func normalizePathSuffix(ul *url.URL) {
	if ul.Path == "" || ul.Path[len(ul.Path)-1] != '/' {
		ul.Path += "/"
	}
}

// normalizeSPHost converts mobile URLs into their PC URLs.
func normalizeSPHost(ul *url.URL) {
	host, ok := spPCHostMap[ul.Host]
	if ok {
		ul.Host = host
	}
}

// normalizePath reduces known URLs to the top page of the website
func normalizePath(ul *url.URL) bool {
	f, ok := normalizePathMap[ul.Host]
	if !ok {
		return false
	}
	return f(ul)
}

var supportedSchemes = [...]string{
	"http",
	"https",
	"mobileapp",
}

func isValidScheme(ul *url.URL) bool {
	for _, s := range supportedSchemes {
		if s == ul.Scheme {
			return true
		}
	}
	return false
}

func isStaticURL(ul *url.URL) bool {
	return ul.Scheme == "mobileapp"
}

// 指定したURLを第N階層で区切り、Domainとします
func splitNDomainPath(ul *url.URL, n int) string {
	path := splitNPath(ul, n)
	if path == "" {
		return ul.Host
	}
	return ul.Host + "/" + path
}

// 指定したPathを第N階層で区切って返します
func splitNPath(ul *url.URL, n int) string {
	vs := strings.Split(ul.Path, "/")
	if len(vs) <= 1 || vs[1] == "" {
		return ""
	}
	parts := []string{}
	size := len(vs)
	if vs[size-1] == "" {
		size--
		vs = vs[:size]
	}
	if vs[0] == "" {
		size--
		vs = vs[1:]
	}
	for i := 0; i < size && i < n; i++ {
		parts = append(parts, vs[i])
	}
	return strings.Join(parts, "/")
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

func makeStringBoolMap(lines []string) map[string]bool {
	m := make(map[string]bool, len(lines))
	for _, line := range lines {
		m[line] = true
	}
	return m
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

func init() {
	var err error

	disusedParametersMap = makeStringBoolMap([]string{
		"fb_action_ids",
		"fb_action_types",
		"fb_source",
		"action_object_map",
		"action_type_map",
		"action_ref_map",
	})
	spPCHostMap = makeStringStringMap(strings.Split(spHostData, "\n"), ",")
	normalizePathMap, err = makeNormalizePathMap(strings.Split(pathDepthData, "\n"), ",")
	if err != nil {
		panic(err)
	}
}

const spHostData = `sp.nicovideo.jp,www.nicovideo.jp
touch.pixiv.net,www.pixiv.net
touch.allabout.co.jp,allabout.co.jp
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
blog.kuruten.jp,1,
blog.livedoor.jp,1,
blog.oricon.co.jp,1,
blogs.yahoo.co.jp,1,
cp.atrct.tv 2,
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
mblg.tv,1
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
