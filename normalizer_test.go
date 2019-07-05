// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package main

import (
	"encoding/json"
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
}

func TestNormalization(t *testing.T) {
	files, err := filepath.Glob(testDir + "/*.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, fn := range files {
		t.Run(fn, func(t *testing.T) {
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
		return nil, err
	}

	tf := testFile{}
	if err := json.NewDecoder(f).Decode(&tf); err != nil {
		return nil, err
	}

	return &tf, nil
}

func testNormalize(t *testing.T, tests []testData) {
	for _, tt := range tests {
		if tt.N1URL == "" && tt.N2URL == "" {
			t.Errorf("%s: n1url or n2url is required", tt.In)
			continue
		}
		ul, err := url.Parse(tt.In)
		if err != nil {
			t.Errorf("%v : %v", tt.In, err)
		}
		t.Run(ul.Host+ul.Path, func(t *testing.T) {
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
		})
	}
}
