// @flow
/*::
type PathDepth = {
  [string]: {
    depth: number,
    paths?: Array<string>,
    query?: Array<string>,
    replace?: {
      pattern: string,
      with: string
    }
  }
}
*/
export var N1URLPathDepthData = {
  'amigo.gesoten.com': {
    depth: 0
  },

  'gaingame.gesoten.com': {
    depth: 0
  },

  'uranai.nosv.org': {
    depth: 1,
    paths: ['/favorite.php']
  },

  'enjoy.point.auone.jp': {
    depth: 1,
    paths: ['/gacha', '/reward', '/enquete']
  },

  's.maho.jp': {
    depth: 3,
    paths: ['/book']
  },

  'ncode.syosetu.com': {
    depth: 2,
    paths: 'ALL'
  },

  'syosetu.org': {
    depth: 3,
    paths: ['/novel']
  },

  'novel.syosetu.org': {
    depth: 2,
    paths: 'ALL'
  },

  'dokuha.jp': {
    depth: 4,
    paths: ['/comicweb'],
    replace: {
      pattern: '/comicweb/viewer',
      with: '/comicweb/contents'
    }
  }
};

export var N2URLPathDepthData = {
  'am-our.com': {
    depth: 1
  },
  'ameblo.jp': {
    depth: 1
  },
  'b.ibbs.info': {
    depth: 1
  },
  'bannch.com': {
    depth: 3
  },
  'bbs.mottoki.com': {
    query: ['bbs'],
    depth: 1
  },
  'bbs7.meiwasuisan.com': {
    depth: 1
  },
  'blog.goo.ne.jp': {
    depth: 1
  },
  'blog.kuruten.jp': {
    depth: 1
  },
  'blog.livedoor.jp': {
    depth: 1
  },
  'blogs.yahoo.co.jp': {
    depth: 1
  },
  'ch.nicovideo.jp': {
    depth: 1
  },
  'cp.atrct.tv': {
    depth: 2
  },
  'd.hatena.ne.jp': {
    depth: 1
  },
  'fanblogs.jp': {
    depth: 1
  },
  'free.jikkyo.org': {
    depth: 4
  },
  'girlsnews.tv': {
    depth: 1
  },
  'ibbs.info': {
    query: ['id'],
    depth: 1
  },
  'jbbs.shitaraba.net': {
    depth: 5
  },
  'lineblog.me': {
    depth: 1
  },
  'lyze.jp': {
    depth: 1
  },
  'mbbs.tv': {
    query: ['id'],
    depth: 1
  },
  'mblg.tv': {
    depth: 1
  },
  'mdpr.jp': {
    depth: 1
  },
  'nanos.jp': {
    depth: 1
  },
  'plaza.rakuten.co.jp': {
    depth: 1
  },
  'rank.log2.jp': {
    depth: 1
  },
  'rara.jp': {
    depth: 1
  },
  's1.ibbs.info': {
    query: ['id'],
    depth: 1
  },
  's2.ibbs.info': {
    query: ['id'],
    depth: 1
  },
  'seesaawiki.jp': {
    depth: 1
  },
  'spora.jp': {
    depth: 1
  },
  'woman.excite.co.jp': {
    depth: 2
  },
  'w.atwiki.jp': {
    depth: 1
  },
  'www.atwiki.jp': {
    depth: 1
  },
  'www.dclog.jp': {
    depth: 1
  },
  'www.ebbs.jp': {
    query: ['b'],
    depth: 1
  },
  'www.eniblo.com': {
    depth: 1
  },
  'www.ne.jp': {
    depth: 3
  },
  'www.nikkan-gendai.com': {
    depth: 3
  },
  'www.plala.or.jp': {
    depth: 1
  },
  'www.tokyo-sports.co.jp': {
    depth: 1
  },
  'www.upp.so-net.ne.jp': {
    depth: 1
  },
  'www.zakzak.co.jp': {
    depth: 1
  },
  'yaplog.jp': {
    depth: 1
  },
  '1st.geocities.jp': {
    depth: 1
  },
  '2nd.geocities.jp': {
    depth: 1
  },
  '3rd.geocities.jp': {
    depth: 1
  },
  'akiba.geocities.jp': {
    depth: 1
  },
  'anime.geocities.jp': {
    depth: 1
  },
  'beauty.geocities.jp': {
    depth: 1
  },
  'book.geocities.jp': {
    depth: 1
  },
  'foodpia.geocities.jp': {
    depth: 1
  },
  'heartland.geocities.jp': {
    depth: 1
  },
  'ichiba.geocities.jp': {
    depth: 1
  },
  'island.geocities.jp': {
    depth: 1
  },
  'milky.geocities.jp': {
    depth: 1
  },
  'motor.geocities.jp': {
    depth: 1
  },
  'movie.geocities.jp': {
    depth: 1
  },
  'music.geocities.jp': {
    depth: 1
  },
  'outdoor.geocities.jp': {
    depth: 1
  },
  'park.geocities.jp': {
    depth: 1
  },
  'sky.geocities.jp': {
    depth: 1
  },
  'space.geocities.jp': {
    depth: 1
  },
  'sports.geocities.jp': {
    depth: 1
  },
  'www.geocities.jp': {
    depth: 1
  },

  'ncode.syosetu.com': {
    depth: 1,
    paths: 'ALL'
  },

  'syosetu.org': {
    depth: 2,
    paths: ['/novel']
  },

  'novel.syosetu.org': {
    depth: 1,
    paths: 'ALL'
  }
};
