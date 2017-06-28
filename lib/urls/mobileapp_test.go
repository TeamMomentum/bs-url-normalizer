package urls

import "testing"

func TestNormalizeMobileApp(t *testing.T) {
	var cases = []struct {
		rawurl string
		wants  string
	}{
		{
			rawurl: "https://itunes.apple.com/jp/app/minkara/id346528801?mt=8",
			wants:  "mobileapp::1-346528801",
		},
		{
			rawurl: "https://itunes.apple.com/app/id346528801",
			wants:  "mobileapp::1-346528801",
		},
		{
			rawurl: "https://play.google.com/store/apps/details?id=net.totopi.news&hl=jp",
			wants:  "mobileapp::2-net.totopi.news",
		},
		{
			rawurl: "https://play.google.com/store/apps/details?id=net.totopi.news",
			wants:  "mobileapp::2-net.totopi.news",
		},
	}

	for _, cs := range cases {
		ul := mustURL(cs.rawurl)
		if u := FirstNormalizeURL(ul); u != cs.wants {
			t.Errorf("%v != %v", u, cs.wants)
		}
	}
}
