// @flow
/*::
type IgnoreQuery = {
  [string]: {
    keys: IgnoreValue,
    paths: IgnoreValue,
  }
}
type IgnoreValue = 'ALL' | Array<string>
*/

export var IgnoreQueryData /*: IgnoreQuery */ = {
  ALL: {
    keys: [
      'fb_action_types',
      'fb_source',
      'action_object_map',
      'action_type_map',
      'action_ref_map',
      'cb',
      'fbclid',
      'mr'
    ],
    paths: 'ALL'
  },

  'amigo.gesoten.com': {
    keys: 'ALL',
    paths: 'ALL'
  },

  'gaingame.gesoten.com': {
    keys: 'ALL',
    paths: 'ALL'
  },

  'uranai.nosv.org': {
    keys: 'ALL',
    paths: ['/favorite.php']
  },

  'live.nicovideo.jp': {
    keys: 'ALL',
    paths: ['/watch']
  },

  'enjoy.point.auone.jp': {
    keys: 'ALL',
    paths: ['/gacha', '/reward', '/enquete']
  },

  'd.pixiv.org': {
    keys: ['num'],
    paths: 'ALL'
  },

  'enq.nstk-4.com': {
    keys: ['time'],
    paths: 'ALL'
  },

  'tabelog.com': {
    keys: 'ALL',
    paths: 'ALL'
  },

  'www.pixiv.net': {
    keys: 'ALL',
    paths: [
      '/member_illust.php',
      '/bookmark',
      '/recommend.php',
      '/novel/recommend.php',
      '/novel/member.php'
    ]
  },

  'www.nicovideo.jp': {
    keys: 'ALL',
    paths: ['/watch/sm']
  },

  'itest.5ch.net': {
    keys: ['url'],
    paths: ['/jump/to']
  },

  'itest.bbspink.com': {
    keys: ['url'],
    paths: ['/jump/to']
  }
};
