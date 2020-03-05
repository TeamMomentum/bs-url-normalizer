// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"reflect"
	"testing"
)

// See also: https://golang.org/src/net/url/url_test.go
// nolint:funlen
func Test_parsePotentialURL(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    *url.URL
		wantErr bool
	}{
		{
			"minimal URL",
			"http://www.example.com",
			&url.URL{
				Scheme: "http",
				Host:   "www.example.com",
			},
			false,
		},
		{
			"ordinal URL",
			"https://www.example.com:8080/path/to/page.html?x=1#99",
			&url.URL{
				Scheme:   "https",
				Host:     "www.example.com:8080",
				Path:     "/path/to/page.html",
				RawQuery: "x=1",
				Fragment: "99",
			},
			false,
		},
		{
			"capital scheme",
			"HTTP://www.example.com",
			&url.URL{
				Scheme: "http",
				Host:   "www.example.com",
			},
			false,
		},
		{
			"non-http but common scheme URL",
			"ftp://ftp.example.com/bar",
			&url.URL{
				Scheme: "ftp",
				Host:   "ftp.example.com",
				Path:   "/bar",
			},
			false,
		},
		{
			"non-http scheme URL with port",
			"ftp://ftp.example.com:10022/bar",
			&url.URL{
				Scheme: "ftp",
				Host:   "ftp.example.com:10022",
				Path:   "/bar",
			},
			false,
		},
		{
			"non-http scheme URL without authority part",
			"ftp://ftp.example.com:10022/bar",
			&url.URL{
				Scheme: "ftp",
				Host:   "ftp.example.com:10022",
				Path:   "/bar",
			},
			false,
		},
		{ // See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URIs
			"non-http uncommon scheme URL",
			"data:text/html,%3Ch1%3EHello%2C%20World!%3C%2Fh1%3E",
			&url.URL{
				Scheme: "data",
				Opaque: "text/html,%3Ch1%3EHello%2C%20World!%3C%2Fh1%3E",
			},
			false,
		},
		{
			"without scheme",
			"www.example.com/foo",
			&url.URL{
				Scheme: "http",
				Host:   "www.example.com",
				Path:   "/foo",
			},
			false,
		},
		{
			"without scheme with port",
			"www.example.com:8080/foo",
			&url.URL{
				Scheme: "http",
				Host:   "www.example.com:8080",
				Path:   "/foo",
			},
			false,
		},
		{
			"multibyte hostname without scheme",
			"hello.世界.com/foo",
			&url.URL{
				Scheme: "http",
				Host:   "hello.世界.com",
				Path:   "/foo",
			},
			false,
		},
		{
			"multibyte hostname with port",
			"http://hello.世界.com:8080/foo",
			&url.URL{
				Scheme: "http",
				Host:   "hello.世界.com:8080",
				Path:   "/foo",
			},
			false,
		},
		{
			"multibyte hostname with port without scheme",
			"hello.世界.com:8080/foo",
			&url.URL{
				Scheme: "http",
				Host:   "hello.世界.com:8080",
				Path:   "/foo",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePotentialURL(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePotentialURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePotentialURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptimizeAMPCacheURL(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *url.URL
	}{
		{
			"simple URL",
			"https://amp-dev.cdn.ampproject.org/c/s/amp.dev/index.amp.html",
			&url.URL{
				Scheme: "https",
				Host:   "amp.dev",
				Path:   "/index.amp.html",
			},
		},
		{
			"URL with parameter and fragment",
			"https://amp-dev.cdn.ampproject.org/c/s/amp.dev/amp/index.html?param=123#Head",
			&url.URL{
				Scheme:   "https",
				Host:     "amp.dev",
				Path:     "/amp/index.html",
				Fragment: "Head",
				RawQuery: "param=123",
			},
		},
		{
			"top page URL (no path)",
			"https://amp-example-com.cdn.ampproject.org/c/amp.example.com",
			&url.URL{
				Scheme: "http",
				Host:   "amp.example.com",
			},
		},
		{
			"empty URL (http://) fails to optimize",
			"https://4oymiquy7qobjgx36tejs35zeqt24qpemsnzgtfeswmrw6csxbkq.cdn.ampproject.org/c/",
			&url.URL{
				Scheme: "https",
				Host:   "4oymiquy7qobjgx36tejs35zeqt24qpemsnzgtfeswmrw6csxbkq.cdn.ampproject.org",
				Path:   "/c/",
			},
		},
		{
			"Japanese domain (ドメイン名例.JP)",
			"https://xn---jp-qi4b8gof5e173yzqi.cdn.ampproject.org/c/s/xn--eckwd4c7cu47r2wf.jp/page/index.amp.html",
			&url.URL{
				Scheme: "https",
				Host:   "xn--eckwd4c7cu47r2wf.jp",
				Path:   "/page/index.amp.html",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.arg)
			got := optimizeAMPCacheURL(u)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optimizeAMPCacheURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
