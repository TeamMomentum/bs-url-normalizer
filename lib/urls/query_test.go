package urls

import (
	"net/url"
	"testing"
)

func Test_removeQueryParameters(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		// TODO: Add test cases.
		{
			in:  "http://optimized-by.rubiconproject.com/a/14312/158276/761644-15.html?&cb=0.30209593979481797&tk_st=1&rp_s=c&p_pos=btf&p_screen_res=360x640&ad_slot=158276_15",
			out: "http://optimized-by.rubiconproject.com/a/14312/158276/761644-15.html?ad_slot=158276_15&p_pos=btf&p_screen_res=360x640&rp_s=c&tk_st=1",
		},
		{
			in:  "https://d.pixiv.org/show?zone_id=t_header&segments=android&format=js&s=0&up=0&ng=w&l=ja&os=and&ngt=w&pla_referer_page_name=pixiv_novel&ab_test_digits_first=8&ab_test_digits_second=6&kw=3162bcf9da3e097&kw=403d9da10783284&num=5a1263d8455",
			out: "https://d.pixiv.org/show?ab_test_digits_first=8&ab_test_digits_second=6&format=js&kw=3162bcf9da3e097&kw=403d9da10783284&l=ja&ng=w&ngt=w&os=and&pla_referer_page_name=pixiv_novel&s=0&segments=android&up=0&zone_id=t_header",
		},
		{
			in:  "http://test.com/a?num=1",
			out: "http://test.com/a?num=1",
		},
		{
			in:  "https://s.tabelog.com/aichi/A2303/A230302/23000859/?svd=20171120&svt=1900&svps=2&default_yoyaku_condition=1",
			out: "https://s.tabelog.com/aichi/A2303/A230302/23000859/",
		},
		{
			in:  "https://touch.pixiv.net/novel/show.php?id=8928225&uarea=tag",
			out: "https://touch.pixiv.net/novel/show.php?id=8928225&uarea=tag",
		},
		{
			in:  "https://touch.pixiv.net/bookmark.php?id=5020681&p=7",
			out: "https://touch.pixiv.net/bookmark.php",
		},
		{
			in:  "https://touch.pixiv.net/novel/recommend.php?id=6160013",
			out: "https://touch.pixiv.net/novel/recommend.php",
		},
		{
			in:  "http://www.nicovideo.jp/search/Trinity%20Field%20MASTER?f_range=0&l_range=0&opt_md=&start=&end=",
			out: "http://www.nicovideo.jp/search/Trinity%20Field%20MASTER?end=&f_range=0&l_range=0&opt_md=&start=",
		},
		{
			in:  "http://www.nicovideo.jp/watch/sm32288578?ref=video%2Franking%2Ffav%2Fdaily%2Fgame%2Fnow%2F50",
			out: "http://www.nicovideo.jp/watch/sm32288578",
		},
	}
	for _, tt := range tests {
		ul, err := url.Parse(tt.in)
		if err != nil {
			t.Errorf("%v : %v", tt.in, err)
		}
		t.Run(ul.Host+ul.Path, func(t *testing.T) {
			removeQueryParameters(ul)
			if ul.String() != tt.out {
				t.Errorf("unexpected %#v : %v", ul, tt.out)
			}
		})
	}
}
