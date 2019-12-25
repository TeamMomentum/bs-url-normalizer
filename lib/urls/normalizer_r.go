// Copyright 2019 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"fmt"
	"net/url"
)

// NormalizerR normalizes the URL into the crawlable URL and the key for KVS use
// Most behavior is as same as `Normalizer`, except `normalizeSPHost` in `FirstNormalizedURL()`
type NormalizerR struct {
	cURL  string
	n1URL string
	n2URL string
	url   *url.URL
}

// NewNormalizerR generate a new NormalizerR structure when the input URL is supported.
// Unlike `NewNormalizer`, argument `ul` will not be destroyed.
func NewNormalizerR(ul *url.URL) (n *NormalizerR, err error) {
	if !isValidScheme(ul) {
		n = nil
		err = fmt.Errorf("invalid scheme %v", ul.Scheme)
		return
	}

	url1 := ""
	url2 := ""
	url3 := ""
	if isStaticURL(ul) {
		url1 = ul.String()
		url2 = url1
		url3 = url1
	}

	// allocate and copy
	newUl := *ul

	n = &NormalizerR{
		url:   &newUl,
		n1URL: url1,
		n2URL: url2,
		cURL:  url3,
	}

	return
}

// CrawlingURL returns the preferred URL for crawling
func (n *NormalizerR) CrawlingURL() string {
	if n.cURL == "" {
		normalizePunycodeHost(n.url)
		n.url = optimizeURL(n.url)
		n.cURL = n.url.String()
	}
	return n.cURL
}

// FirstNormalizedURL returns a unique URL of the input URL,
// which contributes to reduce the database footprint.
func (n *NormalizerR) FirstNormalizedURL() string {
	if n.n1URL != "" {
		return n.n1URL
	}

	n.CrawlingURL()
	ul := n.url

	if mu := normalizeMobileAppURL(ul); mu != "" {
		n.n1URL = mu
		n.n2URL = mu
		return mu
	}

	// temporary apply normalizeSPHost to make removeQueryParameters works correctly
	originalHost := ul.Host
	spNormalized := normalizeSPHost(ul)

	normalizeScheme(ul)
	removeQueryParameters(ul)
	normalizePathSuffix(ul)

	// restore host part if normalizeSPHost was applied
	if spNormalized {
		ul.Host = originalHost
	}

	n.n1URL = ul.String()

	return n.n1URL
}

// SecondNormalizedURL does the FirstNormalizeURL first, then
// shrinks the URL by website as much as possible.
func (n *NormalizerR) SecondNormalizedURL() string {
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
