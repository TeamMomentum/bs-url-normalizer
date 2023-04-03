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
//
//nolint:funlen, exhaustivestruct
func Test_parsePotentialURL(t *testing.T) {
	t.Parallel()

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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
