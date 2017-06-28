// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

func TestIsMobileApp(t *testing.T) {
	type args struct {
		placement string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "MobileApp placement (AppStore)",
			args: args{"mobileapp::1-410395246"},
			want: true,
		}, {
			name: "MobileApp placement (PlayStore)",
			args: args{"mobileapp::2-com.labpixies.colordrips"},
			want: true,
		}, {
			name: "Not MobileApp placement (Invalid format)",
			args: args{"mobileapp::x-com.labpixies.colordrips"},
			want: false,
		}, {
			name: "Not MobileApp placement (Basic URL)",
			args: args{"http://www.example.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMobileApp(tt.args.placement); got != tt.want {
				t.Errorf("IsMobileApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMobileAppPlacement(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name   string
		args   args
		want   *MobileApp
		wantOk bool
	}{
		{
			name:   "MobileApp placement (AppStore)",
			args:   args{"mobileapp::1-410395246"},
			want:   &MobileApp{"mobileapp::1-410395246", "1", "410395246", nil},
			wantOk: true,
		}, {
			name:   "MobileApp placement (PlayStore)",
			args:   args{"mobileapp::2-com.labpixies.colordrips"},
			want:   &MobileApp{"mobileapp::2-com.labpixies.colordrips", "2", "com.labpixies.colordrips", nil},
			wantOk: true,
		}, {
			name:   "MobileApp placement (Unsupported StoreType)",
			args:   args{"mobileapp::0-com.labpixies.colordrips"},
			want:   &MobileApp{"mobileapp::0-com.labpixies.colordrips", "0", "com.labpixies.colordrips", nil},
			wantOk: true,
		}, {
			name:   "MobileApp placement (Empty AppID)",
			args:   args{"mobileapp::0-"},
			want:   nil, // might be parse error
			wantOk: false,
		}, {
			name:   "Not MobileApp placement (Basic URL)",
			args:   args{"http://www.example.com"},
			want:   nil,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := NewMobileAppPlacement(tt.args.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMobileAppPlacement() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("NewMobileAppPlacement() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMobileApp_StorePageURL(t *testing.T) {
	testLocale := language.MustParse("en-US")
	type fields struct {
		RawPlacement string
		StoreType    string
		AppID        string
		Locale       *language.Tag
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:    "MobileApp placement (AppStore, default locale)",
			fields:  fields{"mobileapp::1-410395246", "1", "410395246", nil},
			want:    "https://itunes.apple.com/jp/app/id410395246",
			wantErr: false,
		}, {
			name:    "MobileApp placement (AppStore, with CountryCode)",
			fields:  fields{"mobileapp::1-410395246", "1", "410395246", &testLocale},
			want:    "https://itunes.apple.com/us/app/id410395246",
			wantErr: false,
		},
		{
			name:    "MobileApp placement (PlayStore, default locale)",
			fields:  fields{"mobileapp::2-com.labpixies.colordrips", "2", "com.labpixies.colordrips", nil},
			want:    "https://play.google.com/store/apps/details?id=com.labpixies.colordrips&hl=ja",
			wantErr: false,
		},
		{
			name:    "MobileApp placement (PlayStore, With Lang)",
			fields:  fields{"mobileapp::2-com.labpixies.colordrips", "2", "com.labpixies.colordrips", &testLocale},
			want:    "https://play.google.com/store/apps/details?id=com.labpixies.colordrips&hl=en",
			wantErr: false,
		},
		{
			name:    "MobileApp placement (Unsupported StoreType)",
			fields:  fields{"mobileapp::0-com.labpixies.colordrips", "0", "com.labpixies.colordrips", nil},
			want:    "",
			wantErr: true,
		},
		{
			name:    "MobileApp placement (Empty AppID)",
			fields:  fields{"mobileapp::1-410395246", "1", "", nil}, // would have caused parse error
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := &MobileApp{
				RawPlacement: tt.fields.RawPlacement,
				StoreType:    tt.fields.StoreType,
				AppID:        tt.fields.AppID,
				Locale:       tt.fields.Locale,
			}
			got, err := ma.StorePageURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("MobileApp.StorePageURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MobileApp.StorePageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
