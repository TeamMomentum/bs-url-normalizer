// Copyright 2016 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import "net/url"

// FirstNormalizeURL returns a unique URL of the input URL,
// which contributes to reduce the database footprint.
func FirstNormalizeURL(ul *url.URL) string {
	n, err := NewNormalizer(ul)
	if err != nil {
		return ""
	}

	return n.FirstNormalizedURL()
}

// SecondNormalizeURL does the FirstNormalizeURL first, then
// shrinks the URL by website as much as possible.
func SecondNormalizeURL(ul *url.URL) string {
	n, err := NewNormalizer(ul)
	if err != nil {
		return ""
	}

	return n.SecondNormalizedURL()
}

// CrawlingURL convert URL for crawling.
func CrawlingURL(ul *url.URL) string {
	n, err := NewNormalizer(ul)
	if err != nil {
		return ""
	}

	return n.CrawlingURL()
}

// R1NormalizeURL returns a unique URL of the input URL,
// which contributes to reduce the database footprint.
func R1NormalizeURL(ul *url.URL) string {
	n, err := NewNormalizerR(ul)
	if err != nil {
		return ""
	}

	return n.FirstNormalizedURL()
}

// R2NormalizeURL returns a unique URL of the input URL,
// which contributes to reduce the database footprint.
func R2NormalizeURL(ul *url.URL) string {
	n, err := NewNormalizerR(ul)
	if err != nil {
		return ""
	}

	return n.SecondNormalizedURL()
}
