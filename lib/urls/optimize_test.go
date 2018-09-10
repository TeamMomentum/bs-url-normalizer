// Copyright 2017 Momentum K.K. All rights reserved.
// This source code or any portion thereof must not be
// reproduced or used in any manner whatsoever.

// Package urls normalizes URLs and implements its helpers
package urls

import (
	"net/url"
	"testing"
)

func TestOptimizeURL(t *testing.T) {
	var cases = []struct {
		rawurl string
		wants  string
	}{
		{
			"http://live.nicovideo.jp/watch/lv270002526?ref=notify&zroute=subscribe",
			"http://live.nicovideo.jp/watch/lv270002526",
		},
		{
			"http://ncode.syosetu.com/n7779dh/103/",
			"http://ncode.syosetu.com/n7779dh",
		},
		{
			"http://dokuha.jp/comicweb/viewer/comic/real/1",
			"http://dokuha.jp/comicweb/contents/comic/real",
		},
		{
			"https://novel.syosetu.org/81116/20.html",
			"https://novel.syosetu.org/81116",
		},
		{
			"http://s.maho.jp/book/2f7cc0g0b8fc434d/4767056014/2/",
			"http://s.maho.jp/book/2f7cc0g0b8fc434d/4767056014/",
		},
		{
			"https://enjoy.point.auone.jp/gacha/lottery/?token=+tnmAJSSfYGQ4k9ppHfZPHs2b5OWEYWOzbW1I4VkH",
			"https://enjoy.point.auone.jp/gacha",
		},
		{
			"https://enjoy.point.auone.jp/reward/?medid=walletmail&srcid=tameru&serial=0185&i=AeE9zC&ps=banner",
			"https://enjoy.point.auone.jp/reward",
		},
		{
			"https://enjoy.point.auone.jp/enquete/?aid=guronabi&bid=enquete&cid=",
			"https://enjoy.point.auone.jp/enquete",
		},
		{
			"http://uranai.nosv.org/recommend.php?urid=novel/flato",
			"http://uranai.nosv.org/recommend.php?urid=novel/flato",
		},
		{
			"http://uranai.nosv.org/favorite.php?crumb=0adf265b9d9c7921914c9f9fa32adeb3&add=novel/worldmadeh5&p=33&commu_id=worldmadehappye",
			"http://uranai.nosv.org/favorite.php",
		},
		{
			"http://amigo.gesoten.com/jewel/event/1494614261",
			"http://amigo.gesoten.com",
		},
		{
			"http://gaingame.gesoten.com/gaingame?user_id=000287102299&media_id=56&time=20171121001956&key=15E73E171612396096A3F68D8379841B",
			"http://gaingame.gesoten.com",
		},
		{
			"http://adm.shinobi.jp/a/384b4d631a8cfc69600abf2f7f11a529?x=150&y=0&url=http%3A%2F%2Fwww.example.com%2Fxyz%2Fabc%2Fpage%2F&referrer=&user_id=&du=http%3A%2F%2Fwww.example.com%2Fxyz%2Fabc%2Fpage%2F&iw=300&ih=250",
			"http://www.example.com/xyz/abc/page/",
		},
		{
			"https://googleads.g.doubleclick.net/pagead/ads?client=ca-pub-9033894351963905&output=html&h=250&slotname=1477980675&adk=1268037780&adf=2150039491&w=330&fwrn=2&lmt=1511790354&rafmt=3&format=330x250&url=https%3A%2F%2Fwww.example.com%2F2017%2F11%2F26%2Fpage%2F&region=hajisan&flash=0&fwr=0&resp_fmts=1&sfro=1&wgl=1&adsid=NT&dt=1511790352188&bpp=19&bdt=4339&fdt=2738&idt=2755&shv=r20171113&cbv=r20170110&saldr=aa&prev_fmts=330x60%2C330x250&correlator=8166610681857&frm=20&ga_vid=2067164720.1505658871&ga_sid=1511790355&ga_hid=1111340073&ga_fc=0&pv=1&iag=3&icsg=2&nhd=1&dssz=2&mdo=0&mso=0&u_tz=540&u_his=1&u_java=1&u_h=640&u_w=360&u_ah=640&u_aw=360&u_cd=32&u_nplug=0&u_nmime=0&adx=15&ady=8035&biw=360&bih=540&abxe=1&eid=21061122%2C33895411&oid=3&rx=0&eae=0&fc=668&brdim=0%2C0%2C0%2C0%2C360%2C0%2C360%2C540%2C360%2C540&vis=1&rsz=%7C%7CpeEbr%7C&abl=CS&ppjl=f&pfx=0&fu=144&bc=1&ifi=4&xpc=hdTDSyV9lP&p=https%3A//hajimete-sangokushi.com&dtd=2800",
			"https://www.example.com/2017/11/26/page/",
		},
		{
			"https://securepubads.g.doubleclick.net/gampad/ads?gdfp_req=1&glade_req=1&glv=26&dt=1511772682906&output=html&iu=%2F19153562%2FP_SmartNews&sz=300x250&sfv=1-0-10&correlator=4028602074156422&adk=1314925101&biw=320&bih=568&adx=10&ady=3227.3125&oid=3&u_sd=2&ifi=1&scp=EnableGAS%3DTrue&nhd=1&url=https%3A%2F%2Fwww.example.com%2F2017%2F11%2F27%2F472414%2F&top=www.example.com",
			"https://www.example.com/2017/11/27/472414/",
		},
		{
			"https://pubads.g.doubleclick.net/gampad/ads?os_version=11.4.0&u_sso=p&ios_base_sdk=11.4&request_id=26&u_mwsso=p&u_so=p&js=afma-sdk-i-v7.31.0&is_arr=true&ios_app_volume=1&is_other_audio_playing=0&sai=1&ios_output_volume=0.6875&fbs_aiid=D6558B4EB29A4FE58F87869271F68C58&ios_current_boot_timestamp_ms=1303491339.552&eid=318477469%2C318481073&request_origin=pub&ms=GVFv59jgFzf8vZ2_6_GwP3QMYkZzkR6HtQmN--Avw9b-Qm-UZVwZf_MPcHVs4CFEWyAK_ByHCSLyW_uxt98XFcnhPsk2TaAcsO-gRv3nduZy5MLIQqfPrr_7V6iiq_Lt2nDJnV6zqz600F_SEn3wqfQKhSYghwxekD1Xce2xFWKekR42929xvPmj9Ibwx4c0yqMd9w_bRoQtvWLoFxXIga2ftPxKkd33dOGISnvZYDHLZPUEagCVVmCVYG13JQKBjnxDy2DKKKO4c8Is-jbpsboSdXV7qcDhHjf6eCDne6uwQ31z5Pg8hkEw_H093oRrIa2Im19pl6BzHR_IsNxO3A&hl=ja&u_sd=2&ios_key_window_w=375&cellular_country_code=440&u_w=375&ios_radio=CTRadioAccessTechnologyLTE&u_h=667&submodel=iPhone9%2C1&ios_key_window_h=667&cap_bs=1&net=wi&ios_app_muted=0&binary_arch=arm64&should_silence_audio=0&cellular_network_code=20&_c_csdk_npa_o=false&_ad_b=300x250&format=300x250_as&ad_x=37.5&ad_y=275&ad_w=300&ad_h=250&ad_v=true&_package_name=com.square-enix.mangaupjp&an=43.0.0.iphone.com.square-enix.mangaupjp&u_audio=5&swipeable=1&dtxcb=9F2000&dtsdk=iphoneos11.4&adk=2672299862&preqs=10&seq_num=18&pimp=10&preqs_in_session=10&time_in_session=486.2405869998038&basets=63395&bas_on=0&oar=0&currts=549645&bas_off=0&crqc=1&treq=519272&tfetch=519361&tresponse=519615&tload=520370&dload=1098&output=html&region=mobile_app&u_tz=540&url=43.0.0.iphone.com.square-enix.mangaupjp.adsenseformobileapps.com&gdfp_req=1&markup=html&m_ast=afmajs&impl=ifr&iu=%2F9116787%2F1284627&sz=300x250&correlator=1681027331541702&_hl=ja-jp&gsb=wi&caps=interactiveVideo_inlineVideo_mraid1_mraid2_th_autoplay_mediation_av_sdkAdmobApiForAds_di_transparentBackground_sdkVideo_aso_sfv_dinm_dim_nav_navc_ct_dinmo_gls_saiMacro_gcache&swdr=fals",
			"mobile-app::1-com.square-enix.mangaupjp",
		},
		{
			"https://pubads.g.doubleclick.net/gampad/ads?_activity_context=false&android_num_video_cache_tasks=0&caps=inlineVideo_interactiveVideo_mraid1_mraid2_sdkVideo_th_autoplay_mediation_av_transparentBackground_sdkAdmobApiForAds_di_aso_sfv_dinm_dim_nav_navc_dinmo_ipdof_gls_gcache_xSeconds&eid=318478496%2C318481016%2C318481687&format=320x50_mb&heap_free=1898008&heap_max=536870912&heap_total=48598368&js=afma-sdk-a-v12843999.11717000.1&msid=jp.takke.android.tkmixiviewer&preqs=6&scroll_index=-1&seq_num=7&target_api=16	",
			"mobile-app::2-jp.takke.android.tkmixiviewer",
		},
		{
			"http://d.socdm.com/adsv/v1?posall=SSPLOC&id=16795&tp=http%3A%2F%2Fwww.example.com%2Ftest%2Fabc.def%2Ftoday%2F1466646832&pp=http%3A%2F%2Fwww.example.com%2Ftest%2Fabc.def%2Ftoday%2F1466646832&rnd=1509797344050&targetID=adg_16795&sdkver=1.7.0&sdktype=0&acl=off",
			"http://www.example.com/test/abc.def/today/1466646832",
		},
		{
			"http://d.socdm.com/adsv/v1?posall=SSPLOC&id=11426&sdktype=1&sdkver=1.5.0&appname=マイブックマーク&appbundle=com.ululu.android.apps.my_bookmark&appver=2.9.6.02&advertising_id=aa5621ad-7339-4ef8-938b-bb68f130bc06&lang=ja&locale=ja_JP&tz=Asia/Tokyo",
			"mobile-app::2-com.ululu.android.apps.my_bookmark",
		},
		{
			"http://d.socdm.com/adsv/v1?posall=SSPLOC&id=22122&sdktype=2&sdkver=1.5.2&appname=FreeTube&appbundle=com.satohiro.playtube&appver=1.5.3.2&networktype=carrier&carrier=440-20&lang=ja-JP&locale=ja_JP&tz=Asia/Tokyo&scheme=freetube",
			"mobile-app::1-com.satohiro.playtube",
		},
		{
			"http://showads.pubmatic.com/AdServer/AdServerServlet?pubId=137870&siteId=215541&adId=1163866&kadwidth=320&kadheight=50&SAVersion=2&js=1&kdntuid=1&pageURL=http%3A%2F%2Fwww.example.com%2Fplus%2Fevent%3Futm_source%3D%26utm_medium%3Dbanner%26utm_campaign%3D%25E3%2583%25A2%25E3%2583%25A1%25E3%2583%25B3%25E3%2582%25BF%25E3%2583%25A0&inIframe=1&kadpageurl=http%3A%2F%2Fwww.example.com&operId=1&kltstamp=2017-11-27%208%3A57%3A9&timezone=9&screenResolution=414x736&ranreq=0.39262822122021623&pmUniAdId=0&adVisibility=2&adPosition=1403x0&dspids=%7B%22uids%22%3A%5B%5D%7D",
			"http://www.example.com/plus/event?utm_source=&utm_medium=banner&utm_campaign=%E3%83%A2%E3%83%A1%E3%83%B3%E3%82%BF%E3%83%A0",
		},
		{
			"http://s.yimg.jp/images/listing/tool/yads/yads-iframe.html/?enc=UTF-8&fr_id=yads_4882764-0&fr_support=1&page=1&pv_ts=1511772692009-3522396&s=98335_206734-229039&ssl=1&t=f&tag_path=https%3A%2F%2Fyads.yjtag.yahoo.co.jp%2Ftag&tagpos=0x0&u=https%3A%2F%2Fwww.example.com%2Fxx111%2Fschedule.html&xd_support=1",
			"https://www.example.com/xx111/schedule.html",
		},
		{
			"http://i.yimg.jp/images/listing/tool/yads/yads-iframe.html?s=53959_12054-221718&t=f&ssl=0&fr_id=yads_3895185-0&xd_support=1&fl_support=27&fr_support=1&enc=UTF-8&pv_ts=1511740652699-346455&tag_path=https%3a%2f%2fyads.yjtag.yahoo.co.jp%2ftag&page=1&u=http%3a%2f%2fwww.example.com%2f&ref=http%3a%2f%2fwww.example.com%2f&tagpos=0x0",
			"http://www.example.com/",
		},
		{
			"http://ssl.webtracker.jp/res/?arid=9-171127180106-22015925&cid=adb9b695990bd8d6846e86a4279677d1fcade5d3d6f49b3ba3f5dc0cd6c0c7406453b36ae9cacd78dcede31dd79d7d39&ssl=1&url=https%3A%2F%2Fwww.example.com%2Fnews%2F201711%2F24146736.html",
			"https://www.example.com/news/201711/24146736.html",
		},
		{
			"http://a.t.webtracker.jp/res/?cid=2a7a48984e3409d5b8e3f8a30e53a2fa96eaca76118c8104aae99617d906c66e1d2324a6061a8a43ebe9feae5a00c6b4&euid=2ed53b38d25c38c35d884b17681ada65868d3eb10bafb869&url=http%3A%2F%2Fwww.example.com%2Felem%2F000%2F001%2F593%2F1593104%2F&arid=30-171127090322-650739457",
			"http://www.example.com/elem/000/001/593/1593104/",
		},
		{
			"http://adw.addlv.smt.docomo.ne.jp/tafs/p/sbst/?_adcount=1&_adinf=16119%7C0_&_aid=2792&_creativeids=56842_&_divid=daisy_2261_00001&_fmt=js&_format=1&_frameid=1511772999080917915229&_ftype=1&_nocache=151177299988269510277&_slot=2261&_url=https%3A%2F%2Fwww.example.com%2Fentertainment%2Fcolumn%2F3921",
			"https://www.example.com/entertainment/column/3921",
		},
		{
			"https://ad.deqwas-dsp.net/AdService/collection.aspx?client=ca-pub-9611473549939085&output=html&h=280&slotname=4942526061&adk=4069183175&adf=3349651995&w=435&fwrn=4&fwrnh=100&lmt=1536209925&loeid=201222032&rafmt=1&guci=2.2.0.0.2.2.0&format=435x280&url=https%3A%2F%2Flimo.media%2Farticles%2F-%2F7319&flash=30.0.0&fwr=0&rh=0&rw=435.2&resp_fmts=3&wgl=1&adsid=ADfwbuh-hVGHjtH6ri2p-tQOeMN4jNRxy_VmynyM9WH05QJTZRu4IHxI39UEKOkB9VEN&dt=1536209925079&bpp=4&bdt=688&fdt=269&idt=302&shv=r20180829&cbv=r20180604&saldr=aa&abxe=1&prev_fmts=870x90,435x280&correlator=5677701074903&frm=20&pv=1&ga_vid=1662089437.1531869027&ga_sid=1536209925&ga_hid=576573333&ga_fc=0&icsg=46182217427104&dssz=52&mdo=0&mso=0&u_tz=540&u_his=10&u_java=1&u_h=818&u_w=1455&u_ah=791&u_aw=1455&u_cd=24&u_nplug=1&u_nmime=2&adx=450&ady=3839&biw=1275&bih=703&scr_x=0&scr_y=0&eid=21060853,201222022&oid=3&ref=https%3A%2F%2Flimo.media%2F&rx=0&eae=0&fc=1808&docm=11&brdim=164,73,-7,-7,1455,,1469,805,1291,718&vis=1&rsz=||leEbr|&abl=CS&ppjl=f&pfx=0&fu=1168&bc=1&ifi=14&xpc=ZOFsd15qSw&p=https%3A%2F%2Flimo.media&dtd=346",
			"https://limo.media/articles/-/7319",
		},
		{
			"https://krad20.deqwas.net/AdService/Collection.aspx?s=18272_442-217134&t=f&ssl=1&fr_id=yads_7437856-2&xd_support=1&fr_support=1&sb_support=0&enc=UTF-8&pv_ts=1536209929316-6281701&tag_path=https%3A%2F%2Fyads.yjtag.yahoo.co.jp%2Ftag&page=1&u=https%3A%2F%2Fmatome.naver.jp%2Fm%2Fodai%2F2140585444076756301%3Fpage%3D2&canu=https%3A%2F%2Fmatome.naver.jp%2Fodai%2F2140585444076756301&tagpos=0x3935&async=0",
			"https://matome.naver.jp/m/odai/2140585444076756301?page=2",
		},

		// unoptimizable
		{
			"https://pubads.g.doubleclick.net/gampad/ads?_activity_context=true&android_num_video_cache_tasks=0&caps=inlineVideo_interactiveVideo_mraid1_mraid2_sdkVideo_th_autoplay_mediation_av_transparentBackground_sdkAdmobApiForAds_di_aso_sfv_dinm_dim_dinmo_gcache&eid=318475406%2C318478658%2C318478607&format=320x50_mb&heap_free=7037696&heap_max=134217728&heap_total=40550400&js=afma-sdk-a-v11746038.9452000.2&mr=97878157284952778%2C-115984107731430746%2C6252059087207515892%2C-1297580106298873782%2C-3694031955756562062%",
			"https://pubads.g.doubleclick.net/gampad/ads?_activity_context=true&android_num_video_cache_tasks=0&caps=inlineVideo_interactiveVideo_mraid1_mraid2_sdkVideo_th_autoplay_mediation_av_transparentBackground_sdkAdmobApiForAds_di_aso_sfv_dinm_dim_dinmo_gcache&eid=318475406%2C318478658%2C318478607&format=320x50_mb&heap_free=7037696&heap_max=134217728&heap_total=40550400&js=afma-sdk-a-v11746038.9452000.2&mr=97878157284952778%2C-115984107731430746%2C6252059087207515892%2C-1297580106298873782%2C-3694031955756562062%",
		},
	}

	for _, cs := range cases {
		up, err := url.Parse(cs.rawurl)
		if err != nil {
			t.Error(err)
		}
		opted := optimizeURL(up)
		if opted.String() != cs.wants {
			t.Errorf("%v != %v", opted.String(), cs.wants)
		}
	}
}
