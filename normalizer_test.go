// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/TeamMomentum/bs-url-normalizer/lib/urls"
)

const testDir = "testdata"

type testFile struct {
	Tests []testData `json:"tests"`
}

type testData struct {
	In    string `json:"in"`
	N1URL string `json:"n1url"`
	N2URL string `json:"n2url"`
	CURL  string `json:"curl"`
	R1URL string `json:"r1url"`
	R2URL string `json:"r2url"`
}

func TestNormalization(t *testing.T) {
	t.Parallel()

	files, err := filepath.Glob(testDir + "/*.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, fn := range files {
		fn := fn

		t.Run(fn, func(t *testing.T) {
			t.Parallel()

			tf, err := parseTestFile(fn)
			if err != nil {
				t.Fatalf("%s: %s", fn, err)
			}

			testNormalize(t, tf.Tests)
		})
	}
}

func parseTestFile(fn string) (*testFile, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}

	var tf testFile
	if err := json.NewDecoder(f).Decode(&tf); err != nil {
		return nil, fmt.Errorf("(*json.Decoder).Decode: %w", err)
	}

	return &tf, nil
}

//nolint:gocognit, cyclop
func testNormalize(t *testing.T, tests []testData) {
	t.Helper()

	for _, tt := range tests {
		if tt.CURL == "" && tt.N1URL == "" && tt.N2URL == "" && tt.R1URL == "" && tt.R2URL == "" {
			t.Errorf("%s: curl, n1url or n2url is required", tt.In)

			continue
		}

		ul, err := url.Parse(tt.In)
		if err != nil {
			t.Errorf("%v : %v", tt.In, err)
		}

		t.Run(ul.Host+ul.Path, func(t *testing.T) {
			if tt.CURL != "" {
				result := urls.CrawlingURL(ul)
				if result != tt.CURL {
					t.Errorf("curl:\nwant '%v'\n got '%v'", tt.CURL, result)
				}
			}

			if tt.N1URL != "" {
				result := urls.FirstNormalizeURL(ul)
				if result != tt.N1URL {
					t.Errorf("n1url:\nwant '%v'\n got '%v'", tt.N1URL, result)
				}
			}

			if tt.N2URL != "" {
				result := urls.SecondNormalizeURL(ul)
				if result != tt.N2URL {
					t.Errorf("n2url:\nwant '%v'\n got '%v'", tt.N2URL, result)
				}
			}

			if tt.R1URL != "" {
				result := urls.R1NormalizeURL(ul)
				if result != tt.R1URL {
					t.Errorf("r1url:\nwant '%v'\n got '%v'", tt.R1URL, result)
				}
			}

			if tt.R2URL != "" {
				result := urls.R2NormalizeURL(ul)
				if result != tt.R2URL {
					t.Errorf("r1url:\nwant '%v'\n got '%v'", tt.R2URL, result)
				}
			}
		})
	}
}
