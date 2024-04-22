// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

package urls

import (
	"testing"
)

//nolint:funlen
func Test_normalizePath(t *testing.T) {
	t.Parallel()

	testPaths := []string{
		"example1.com,1,",
		"example2.com,2,",
		"www.example3.com,1,",
		"example4.com,1,q",
	}

	var err error

	normalizePathMap, err = makeNormalizePathMap(testPaths, ",")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		in   string
		want string
	}{
		{
			// not normalize
			"https://www.example.com/path/to/html?q=query&w=word",
			"https://www.example.com/path/to/html?q=query&w=word",
		},
		{
			"https://example1.com/path/to/html?q=query&w=word",
			"https://example1.com/path",
		},
		{
			"https://example2.com/path/to/html?q=query&w=word",
			"https://example2.com/path/to",
		},
		{
			"https://www01.example3.com/path/to/html?q=query&w=word",
			"https://www01.example3.com/path",
		},
		{
			"https://example4.com/path/to/html?q=query&w=word",
			"https://example4.com/path?q=query",
		},
		{
			"https://www.example.com/~user/path/to/html?q=query&w=word",
			"https://www.example.com/~user",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			ul := mustURL(tt.in)
			shouldNormalize := tt.in != tt.want
			isNormalized := normalizePath(ul)

			if isNormalized != shouldNormalize {
				t.Errorf("normalizePath() returns %v, want %v", isNormalized, shouldNormalize)
			}

			if nURL := ul.String(); nURL != tt.want {
				t.Errorf("normalized URL = %v, want %v", nURL, tt.want)
			}
		})
	}
}

func Test_normalizeSPHost(t *testing.T) {
	// initialize spPCHostMap before calling t.Parallel() to avoid data race.
	testHosts := []string{
		"sp.example1.com,www.example1.com",
		"sp.example2.com,example2.com",
	}

	spPCHostMap = makeStringStringMap(testHosts, ",")

	t.Parallel()

	tests := []struct {
		in   string
		want string
	}{
		{
			"https://www.example.com/path/to/html?q=query&w=word",
			"https://www.example.com/path/to/html?q=query&w=word",
		},
		{
			"https://sp.example1.com/path/to/html?q=query&w=word",
			"https://www.example1.com/path/to/html?q=query&w=word",
		},
		{
			"https://sp.example2.com/path/to/html?q=query&w=word",
			"https://example2.com/path/to/html?q=query&w=word",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			ul := mustURL(tt.in)
			normalizeSPHost(ul)

			if nURL := ul.String(); nURL != tt.want {
				t.Errorf("normalized URL = %v, want %v", nURL, tt.want)
			}
		})
	}
}
