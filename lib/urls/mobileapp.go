// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/text/language"
)

var (
	// appstore パターン
	appStorePatterns = []*regexp.Regexp{
		regexp.MustCompile("/app/id([^/]+)"),
		regexp.MustCompile("/app/[^/]+/id([^/]+)"),
	}
)

// normalizeMobileAppURL normalizes mobile appstore URLs
func normalizeMobileAppURL(ul *url.URL) string {
	switch ul.Host {
	case "itunes.apple.com": // for apple store
		return normalizeAppStore(ul)
	case "play.google.com": // for google store
		return normalizePlayStore(ul)
	}
	return ""
}

// normalizeAppStore normalizes apple's app store to mobileapp::1-id
func normalizeAppStore(ul *url.URL) string {
	for _, pattern := range appStorePatterns {
		matches := pattern.FindStringSubmatch(ul.Path)
		if len(matches) > 1 {
			return "mobileapp::1-" + matches[1]
		}
	}
	return ""
}

// normalizePlayStore normalizes google's playstore to mobileapp::2-id
func normalizePlayStore(ul *url.URL) string {
	q := ul.Query()
	id, ok := q["id"]
	if ok {
		return "mobileapp::2-" + id[0]
	}
	return ""
}

var (
	// DefaultLocale is for specifying country in URL of the AppStore website.
	DefaultLocale language.Tag

	// AdWordsMobileAppPattern represents placement pattern in AdWords context.
	// refer to: https://support.google.com/adwords/answer/2454012
	AdWordsMobileAppPattern *regexp.Regexp

	// MobileAppPatterns is the list of MobileApp Placement Patterns.
	// If mathes any of the pattern, it will be considered MobileApp Placement string.
	MobileAppPatterns []*regexp.Regexp

	// appStoreURLFormat represents MobileApp's Web page in iTunes App Store
	appStoreURLFormat = "https://itunes.apple.com/%sapp/id%s"

	// playStoreURLFormat represents MobileApp's Web page in Google Play Store
	playStoreURLFormat = "https://play.google.com/store/apps/details?id=%s&hl=%s"

	// ErrEmptyAppID is error when converting MobileApp placement to PageURL
	ErrEmptyAppID = errors.New("not specified or empty AppID")
)

func init() {
	DefaultLocale = language.MustParse("ja-JP")
	AdWordsMobileAppPattern = regexp.MustCompile(`^mobileapp::(\d+)-(\S+)$`)
	MobileAppPatterns = []*regexp.Regexp{AdWordsMobileAppPattern}
}

// MobileApp represents MobileApp Placement (URL) in OpenRTB
type MobileApp struct {
	RawPlacement string
	StoreType    string
	AppID        string

	// [Optional] represent default locale (lang and country) if empty.
	Locale *language.Tag
}

// Language returns Base Language Code of the MobileApp.Locale.
// If Locale not specified (== nil) in the MobileApp, use DefaultLocale.
// e.g. `ja-JP` => `ja`, `en-US` => `en`
func (ma *MobileApp) Language() string {
	var lang language.Base
	if ma.Locale != nil {
		lang, _ = ma.Locale.Base()
	} else {
		lang, _ = DefaultLocale.Base()
	}
	return lang.String()
}

// CountryCode returns Country Code (Territory Code) of the MobileApp.Locale.
// If Locale not specified (== nil) in the MobileApp, use DefaultLocale.
// e.g. `ja-JP` => `JP`, `en-US` => `US`
func (ma *MobileApp) CountryCode() string {
	var region language.Region
	if ma.Locale != nil {
		region, _ = ma.Locale.Region()
	} else {
		region, _ = DefaultLocale.Region()
	}
	return region.String()
}

// StorePageURL reforms MobileApp placement to it's page in the Application Store website.
func (ma *MobileApp) StorePageURL() (string, error) {
	if ma.AppID == "" {
		return "", ErrEmptyAppID
	}
	var (
		err error
		url string
	)
	switch ma.StoreType {
	case "1":
		url = getAppStorePageURL(ma.AppID, ma.CountryCode())
	case "2":
		url = getPlayStorePageURL(ma.AppID, ma.Language())
	default:
		// iTunes AppStore and Google PlayStore are only supported for now.
		err = fmt.Errorf("unsupported StoreType: %v", ma.StoreType)
	}
	if err != nil {
		return "", err
	}
	return url, nil
}

// IsMobileApp check if string match
func IsMobileApp(placement string) bool {
	for _, pat := range MobileAppPatterns {
		if pat.MatchString(placement) {
			return true
		}
	}
	return false
}

// NewMobileAppPlacement creates MobileApp struct from raw placement string.
// returns (nil, false) when failed to convert.
func NewMobileAppPlacement(p string) (_ *MobileApp, ok bool) {
	// AdWords pattern is only supported for now
	return getAdWordsMobileAppPlacement(p)
}

func getAdWordsMobileAppPlacement(p string) (_ *MobileApp, ok bool) {
	match := AdWordsMobileAppPattern.FindStringSubmatch(p)
	if len(match) != 3 {
		return nil, false
	}
	return &MobileApp{
		RawPlacement: p,
		StoreType:    match[1],
		AppID:        match[2],
	}, true
}

func getAppStorePageURL(appID, countryCode string) string {
	countryPath := strings.ToLower(countryCode)
	if countryPath != "" && countryPath[len(countryPath)-1] != '/' {
		countryPath = countryPath + "/"
	}
	return fmt.Sprintf(appStoreURLFormat, countryPath, appID)
}

func getPlayStorePageURL(appID, language string) string {
	return fmt.Sprintf(playStoreURLFormat, appID, strings.ToLower(language))
}
