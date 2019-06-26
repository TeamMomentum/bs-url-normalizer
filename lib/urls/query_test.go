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
			in:  "http://www.nicovideo.jp/search/Trinity%20Field%20MASTER?f_range=0&l_range=0&opt_md=&start=&end=",
			out: "http://www.nicovideo.jp/search/Trinity%20Field%20MASTER?end=&f_range=0&l_range=0&opt_md=&start=",
		},
		{
			in:  "http://www.nicovideo.jp/watch/sm32288578?ref=video%2Franking%2Ffav%2Fdaily%2Fgame%2Fnow%2F50",
			out: "http://www.nicovideo.jp/watch/sm32288578",
		},
		{
			in:  "http://enq.nstk-4.com/b41c09cb00cb0a09/iframe.php?bnr=geniee&pos=l1&banner_name=WFcriteo&device=pc&time=1511182658378",
			out: "http://enq.nstk-4.com/b41c09cb00cb0a09/iframe.php?banner_name=WFcriteo&bnr=geniee&device=pc&pos=l1",
		},
		{
			in:  "https://www.ranker.com/list/the-7-most-horrifying-things-found-living-inside-humans/beau-iverson?utm_source=facebook&utm_medium=creepy&pgid=1011190218967434&utm_campaign=squid-babies-image&fbclid=IwAR3tuHm8oTsHs8aDx8ZHHbAPg5-NCgFIqAzJcYto2wngZe0Ylc_6k3q2JFU",
			out: "https://www.ranker.com/list/the-7-most-horrifying-things-found-living-inside-humans/beau-iverson?pgid=1011190218967434",
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
