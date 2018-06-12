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

// 不要なパラメータを除去できているかのtest
func TestRemoveQueryParameters(t *testing.T) {
	var (
		ul *url.URL
		nu string
	)

	ul = mustURL("http://blog.example.jp/tihoukoumu?utm_source=yahoo&utm_medium=cpc&utm_campaign=momentum&key=value")
	nu = FirstNormalizeURL(ul)
	stringCheck(t, "URL", "http://blog.example.jp/tihoukoumu/?key=value", nu)
}

/*
Query parameterの順序保証が保たれているかのtest
*元URL
http://example.com/tihoukoumu?d=1&a=2&c=3&b=4

*正規化後URL
http://example.com/tihoukoumu?a=2&b=4&c=3&d=1
*/
func TestQueryOrder(t *testing.T) {
	testURL := "http://example.com/tihoukoumu?d=1&a=2&c=3&b=4"
	results := make(map[string]bool)
	for i := 0; i < 100; i++ {
		nu := FirstNormalizeURL(mustURL(testURL))
		results[nu] = true
	}
	if len(results) != 1 {
		t.Error("URL query order should be stable.")
	}
	testURL = "http://example.com/tihoukoumu?a=2&c=3&b=4&d=1&utm_query=1"
	for i := 0; i < 100; i++ {
		nu := FirstNormalizeURL(mustURL(testURL))
		results[nu] = true
	}
	if len(results) != 1 {
		t.Error("URL query order should be stable.")
	}
}

// パス階層分割関数のtest
func TestSplitNDomainPath(t *testing.T) {
	ul := mustURL("http://example.com/a/b/c/")
	if splitNDomainPath(ul, 2) != "example.com/a/b" {
		t.Errorf("%v != %v", splitNDomainPath(ul, 2), "example.com/a/b")
	}
}

func stringCheck(t *testing.T, key, correct, other string) {
	if correct != other {
		t.Errorf("%v should be '%v', not '%v'.", key, correct, other)
	}
}

func mustURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}

func TestNewNormalizer(t *testing.T) {
	type args struct {
		ul *url.URL
	}
	hogehoge, _ := url.Parse("hogehoge:://test")
	tests := []struct {
		name    string
		args    args
		wantN   *Normalizer
		wantErr bool
	}{
		{
			name:    "Unsupported scheme",
			args:    args{ul: hogehoge},
			wantN:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := NewNormalizer(tt.args.ul)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNormalizer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNormalizer() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNormalizer_CrawlingURL(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "mobileapp",
			raw:  "mobileapp::1-123456789",
			want: "mobileapp::1-123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul, _ := url.Parse(tt.raw)
			n, err := NewNormalizer(ul)
			if err != nil {
				t.Errorf("Unsupported URL %v", ul.String())
			}
			if got := n.CrawlingURL(); got != tt.want {
				t.Errorf("Normalizer.CrawlingURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizer_Punyclde(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		// examples from https://ja.wikipedia.org/wiki/Punycode
		{
			name: "ドメイン名例",
			raw:  "http://ドメイン名例.jp",
			want: "http://xn--eckwd4c7cu47r2wf.jp",
		},
		{
			name: "ウィキペディア.ドメイン名例",
			raw:  "http://ウィキペディア.ドメイン名例.jp",
			want: "http://xn--cckbak0byl6e.xn--eckwd4c7cu47r2wf.jp",
		},
		{
			name: "例え.テスト",
			raw:  "http://例え.テスト",
			want: "http://xn--r8jz45g.xn--zckzah",
		},

		{
			name: "ドメイン名例 escaped",
			raw:  "http://%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E5%90%8D%E4%BE%8B.jp",
			want: "http://xn--eckwd4c7cu47r2wf.jp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul, _ := url.Parse(tt.raw)
			n, err := NewNormalizer(ul)
			if err != nil {
				t.Errorf("Unsupported URL %v", ul.String())
			}
			if got := n.CrawlingURL(); got != tt.want {
				t.Errorf("Normalizer.CrawlingURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
