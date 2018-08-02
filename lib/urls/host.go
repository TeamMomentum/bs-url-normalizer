// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/TeamMomentum/bs-url-normalizer/lib/assets"
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
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		rows := strings.Split(trimmed, sep)
		if len(rows) < 2 {
			continue
		}
		spHost := rows[0]
		pcHost := rows[1]
		m[spHost] = pcHost
	}
	return m
}

// makeNormalizePathMap: 指定したdomain => N階層のpairからパス正規化パターンを生成します
// - format: `domain,パス階層,残したいparamID` の順に sep 区切り
func makeNormalizePathMap(lines []string, sep string) (map[string]func(*url.URL) bool, error) {
	m := make(map[string]func(*url.URL) bool)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		rows := strings.Split(trimmed, sep)
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
			ul.Fragment = "" // 第一正規化でも実施される可能性があるが、念のため
			return true
		}
	}
	return m, nil
}

var (
	spHostData, pathDepthData string
)

func init() {
	var err error
	spHostData = string(assets.MustAsset("norm_host_sp.csv"))
	pathDepthData = string(assets.MustAsset("norm_host_path.csv"))
	spPCHostMap = makeStringStringMap(strings.Split(spHostData, "\n"), ",")
	normalizePathMap, err = makeNormalizePathMap(strings.Split(pathDepthData, "\n"), ",")
	if err != nil {
		panic(err)
	}
}
