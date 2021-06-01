// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"strings"
)

var (
	disusedParametersMap    map[string]bool
	disusedHostParameterMap map[string]string
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
	hostKey, isDisusedHost := disusedHostParameterMap[ul.Host]
	for key := range ul.Query() {
		if strings.HasPrefix(key, "utm_") {
			query.Del(key)
		} else if isDisusedHost && key == hostKey {
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
		"fb_action_ids",
		"fb_action_types",
		"fb_source",
		"action_object_map",
		"action_type_map",
		"action_ref_map",
		"cb",
		"fbclid",
		"_gl", // https://github.com/TeamMomentum/bs-url-normalizer/issues/71
	})

	disusedHostParameterMap = map[string]string{
		"d.pixiv.org":    "num",
		"enq.nstk-4.com": "time",
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
