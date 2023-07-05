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
      'cb',
      // Facebook
      'fbclid',
      // Google
      '_gl',
      'gclid',
      // 'utm_xxxx', // parameters with 'utm_' prefix are removed without using this array
      'dclid',
      'wbraid',
      'gbraid',
      // HubSpot
      '_hsenc',
      '_hsmi',
      // Marketo
      'mkt_tok',
      // Microsoft
      'cvid',
      'ocid',
      'msclkid',
      // Twitter
      'twclid',
      // Yahoo! JAPAN
      'yclid',
    ],
    paths: 'ALL',
  },

  'amigo.gesoten.com': {
    keys: 'ALL',
    paths: 'ALL',
  },

  'gaingame.gesoten.com': {
    keys: 'ALL',
    paths: 'ALL',
  },

  'uranai.nosv.org': {
    keys: 'ALL',
    paths: ['/favorite.php'],
  },

  'live.nicovideo.jp': {
    keys: 'ALL',
    paths: ['/watch'],
  },

  'enjoy.point.auone.jp': {
    keys: 'ALL',
    paths: ['/gacha', '/reward', '/enquete'],
  },

  'd.pixiv.org': {
    keys: ['num'],
    paths: 'ALL',
  },

  'enq.nstk-4.com': {
    keys: ['time'],
    paths: 'ALL',
  },

  'tabelog.com': {
    keys: 'ALL',
    paths: 'ALL',
  },

  'www.pixiv.net': {
    keys: 'ALL',
    paths: [
      '/member_illust.php',
      '/bookmark',
      '/recommend.php',
      '/novel/recommend.php',
      '/novel/member.php',
    ],
  },

  'www.nicovideo.jp': {
    keys: 'ALL',
    paths: ['/watch/sm'],
  },

  'itest.5ch.net': {
    keys: ['url'],
    paths: ['/jump/to'],
  },

  'itest.bbspink.com': {
    keys: ['url'],
    paths: ['/jump/to'],
  },
};
