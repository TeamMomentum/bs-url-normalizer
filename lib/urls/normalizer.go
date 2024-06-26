// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	defaultScheme = "http"
)

// Normalizer normalizes the URL into the crawlable URL and the key for KVS use.
type Normalizer struct {
	cURL  string
	n1URL string
	n2URL string
	url   *url.URL
}

var ErrInvalidScheme = errors.New("invalid scheme")

// NewNormalizer generate a new Normalizer structure when the input URL is supported.
func NewNormalizer(ul *url.URL) (*Normalizer, error) {
	if !isValidScheme(ul) {
		return nil, fmt.Errorf("%w: %v", ErrInvalidScheme, ul.Scheme)
	}

	url1 := ""
	url2 := ""
	url3 := ""

	if isStaticURL(ul) {
		url1 = ul.String()
		url2 = url1
		url3 = url1
	}

	return &Normalizer{
		url:   ul,
		n1URL: url1,
		n2URL: url2,
		cURL:  url3,
	}, nil
}

// CrawlingURL returns the preferred URL for crawling.
func (n *Normalizer) CrawlingURL() string {
	if n.cURL == "" {
		normalizeHost(n.url)
		n.url = optimizeURL(n.url)
		n.cURL = n.url.String()
	}

	return n.cURL
}

// FirstNormalizedURL returns a unique URL of the input URL,
// which contributes to reduce the database footprint.
func (n *Normalizer) FirstNormalizedURL() string {
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

	normalizeSPHost(ul)
	normalizeScheme(ul)
	removeQueryParameters(ul)
	normalizePathSuffix(ul)
	n.n1URL = ul.String()

	return n.n1URL
}

// SecondNormalizedURL does the FirstNormalizeURL first, then
// shrinks the URL by website as much as possible.
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

// normalizeScheme replaces all schemes into http.
func normalizeScheme(ul *url.URL) {
	if isStaticURL(ul) {
		return
	}

	ul.Scheme = defaultScheme
}

// normalizePathSuffix keeps the end of URL as '/'.
func normalizePathSuffix(ul *url.URL) {
	if ul.Path == "" || ul.Path[len(ul.Path)-1] != '/' {
		ul.Path += "/"
	}
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

// 指定したPathを第N階層で区切って返します.
func splitNPath(ul *url.URL, n int) string {
	vs := strings.Split(ul.Path, "/")
	if len(vs) <= 1 || vs[1] == "" {
		return ""
	}

	size := len(vs)
	if vs[size-1] == "" {
		size--

		vs = vs[:size]
	}

	if vs[0] == "" {
		size--

		vs = vs[1:]
	}

	parts := []string{}

	for i := 0; i < size && i < n; i++ {
		parts = append(parts, vs[i])
	}

	return strings.Join(parts, "/")
}
