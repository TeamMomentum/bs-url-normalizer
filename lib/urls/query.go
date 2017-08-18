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
	disusedParametersMap map[string]bool
)

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
	})
}
