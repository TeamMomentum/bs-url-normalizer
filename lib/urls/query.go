// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"slices"
	"strings"
)

var (
	disusedParametersMap    map[string]bool
	disusedHostParameterMap map[string][]string
	noQueryHostPathsMap     map[string][]string
)

// removeQueryParameters removes the unnecessary query parameters.
func removeQueryParameters(ul *url.URL) {
	if paths, ok := noQueryHostPathsMap[ul.Host]; ok {
		for _, key := range paths {
			if strings.HasPrefix(ul.Path, key) {
				ul.RawQuery = ""

				return
			}
		}
	}

	query := ul.Query()
	hostKeys, isDisusedHost := disusedHostParameterMap[ul.Host]

	for key := range ul.Query() {
		if strings.HasPrefix(key, "utm_") {
			query.Del(key)
		} else if isDisusedHost && slices.Contains(hostKeys, key) {
			query.Del(key)
		} else if _, ok := disusedParametersMap[key]; ok {
			query.Del(key)
		}
	}

	ul.RawQuery = query.Encode()
}

func makeStringBoolMap(lines []string) map[string]bool {
	m := make(map[string]bool, len(lines))
	for _, line := range lines {
		m[line] = true
	}

	return m
}

func init() {
	disusedParametersMap = makeStringBoolMap([]string{
		"cb",
		// Facebook
		"fbclid",
		// Google
		"_gl", // https://github.com/TeamMomentum/bs-url-normalizer/issues/71
		"gclid",
		// "utm_xxxx", // parameters with "utm_" prefix are removed without using this map
		"dclid",
		"wbraid",
		"gbraid",
		// HubSpot
		"_hsenc",
		"_hsmi",
		// Marketo
		"mkt_tok",
		// Microsoft
		"cvid",
		"ocid",
		"msclkid",
		// Twitter
		"twclid",
		// Yahoo! JAPAN
		"yclid",
	})

	disusedHostParameterMap = map[string][]string{
		"d.pixiv.org":    {"num"},
		"enq.nstk-4.com": {"time"},
		"www.msn.com":    {"ei", "pc"},
	}

	noQueryHostPathsMap = map[string][]string{
		"tabelog.com": {""},
		"www.pixiv.net": {
			"/member_illust.php",
			"/bookmark",
			"/recommend.php",
			"/novel/recommend.php",
			"/novel/member.php",
		},
		"www.nicovideo.jp":  {"/watch/sm"},
		"itest.5ch.net":     {"/jump/to"},
		"itest.bbspink.com": {"/jump/to"},
	}
}
