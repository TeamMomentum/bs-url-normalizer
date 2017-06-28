package urls

import (
	"net/url"
	"regexp"
)

var (
	// appstore パターン
	appStorePatterns = []*regexp.Regexp{
		regexp.MustCompile("/app/id([^/]+)"),
		regexp.MustCompile("/app/[^/]+/id([^/]+)"),
	}
)

// normalizeMobileAppURL normalizes mobile appstore URLs
func normalizeMobileAppURL(ul *url.URL) string {
	switch ul.Host {
	case "itunes.apple.com": // for apple store
		return normalizeAppStore(ul)
	case "play.google.com": // for google store
		return normalizePlayStore(ul)
	}
	return ""
}

// normalizeAppStore normalizes apple's app store to mobileapp::1-id
func normalizeAppStore(ul *url.URL) string {
	for _, pattern := range appStorePatterns {
		matches := pattern.FindStringSubmatch(ul.Path)
		if len(matches) > 1 {
			return "mobileapp::1-" + matches[1]
		}
	}
	return ""
}

// normalizePlayStore normalizes google's playstore to mobileapp::2-id
func normalizePlayStore(ul *url.URL) string {
	q := ul.Query()
	id, ok := q["id"]
	if ok {
		return "mobileapp::2-" + id[0]
	}
	return ""
}
