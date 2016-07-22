package urls

import (
	"net/url"
	"regexp"
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
	// appstore パターン
	appStorePatterns = []*regexp.Regexp{
		regexp.MustCompile("/app/(id[^/]+)"),
		regexp.MustCompile("/app/[^/]+/(id[^/]+)"),
	}
)

/*
1段階目評価時関数:
* 特定のクエリパラメータを除去する
* PC・モバイルのページ統合を行う
* http/https プロトコルのhttpで統一する
* パス末尾の正規化
*/
func FirstNormalizeURL(ul *url.URL) string {
	removeQueryParameters(ul, ul.Query())
	normalizeSPHost(ul)
	normalizeScheme(ul)
	normalizeMobileAppURL(ul)
	normalizePathSuffix(ul)
	return ul.String()
}

/*
2段階目評価時関数:
1段階目評価時関数に加え以下の処理
* ドメイン(パス)正規化処理
if 正規化対象のドメインであれば:
  正規化済みドメインを返す
else:
  ドメインのみを返す
*/
func SecondNormalizeURL(ul *url.URL) string {
	FirstNormalizeURL(ul)
	if normalizePath(ul) {
		return ul.String()
	}
	return ul.Scheme + "://" + ul.Host
}

// 不要なクエリパラメータを除去します
func removeQueryParameters(ul *url.URL, query url.Values) {
	deleteQueries := []string{}
	for key, _ := range query {
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

// http/https を常にhttpを利用するように変更します
func normalizeScheme(ul *url.URL) {
	ul.Scheme = defaultScheme
}

// URLパスの末尾が常に'/'になるようにします
func normalizePathSuffix(ul *url.URL) {
	if ul.Path == "" || ul.Path[len(ul.Path)-1] != '/' {
		ul.Path += "/"
	}
}

// SPのドメインをPCのドメインに変換します
// ul: 変換対象URLへの参照
func normalizeSPHost(ul *url.URL) {
	host, ok := spPCHostMap[ul.Host]
	if ok {
		ul.Host = host
	}
}

// パス階層での正規化を行います
// ul: 変換対象URLへの参照
// return: 正規化を行ったかどうかの真偽値
func normalizePath(ul *url.URL) bool {
	f, ok := normalizePathMap[ul.Host]
	if !ok {
		return false
	}
	return f(ul)
}

// mobile appstoreの正規化を行います
func normalizeMobileAppURL(ul *url.URL) {
	switch ul.Host {
	case "itunes.apple.com": // for apple store
		normalizeAppStore(ul)
	}
}

// apple app storeの正規化を行います
// Patterns:
// https://itunes.apple.com/jp/app/minkara/id346528801?mt=8
// https://itunes.apple.com/app/id994362719
func normalizeAppStore(ul *url.URL) {
	for _, pattern := range appStorePatterns {
		matches := pattern.FindStringSubmatch(ul.Path)
		if len(matches) > 1 {
			ul.Path = "/app/" + matches[1]
			ul.RawQuery = ""
			return
		}
	}
}

// 意味空間でURLを切り上げ、crawling対象のURLに変換する関数を返します
func CreateOptimizeURLFunc(re *regexp.Regexp) func(*url.URL) bool {
	return func(ul *url.URL) bool {
		groups := re.FindStringSubmatch(ul.Path)
		if len(groups) == 0 {
			return false
		}
		ul.Path = groups[0]
		return true
	}
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
		size -= 1
		vs = vs[:size]
	}
	if vs[0] == "" {
		size -= 1
		vs = vs[1:]
	}
	for i := 0; i < size && i < n; i++ {
		parts = append(parts, vs[i])
	}
	return strings.Join(parts, "/")
}

// 指定したURLを第N階層で区切り、Domainとします
func splitNDomainPath(ul *url.URL, n int) string {
	path := splitNPath(ul, n)
	if path == "" {
		return ul.Host
	}
	return ul.Host + "/" + path
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
				for key, _ := range query {
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
