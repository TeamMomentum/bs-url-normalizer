// @flow
// refer ../../lib/urls/optimize.go
/*::
type AdFrameParam =
  | {|
      url: string,
    |}
  | {|
      url: string,
      ref: string,
    |}
  | {|
      domain: string,
      url: string,
    |}
  | {|
      domain: string,
      url: string,
      ref: string,
    |}
  | {|
      url: string,
      ref: string,
      android: string,
      ios: string,
    |}
  | {|
      url: string,
      ref: string,
      android: string,
      ios: string,
      content_url: string,
    |}
  | {|
      domain: string,
      url: string,
      ref: string,
      android: string,
      ios: string,
      content_url: string,
    |};
*/

export var adFrameParams /*: { [string]: AdFrameParam } */ = {
  // for verification
  'm0mentum-tags.s3.amazonaws.com': {
    domain: 'domain',
    url: 'url',
    ref: 'ref',
    android: 'an',
    ios: 'ios',
    content_url: 'content_url',
  },

  'adm.shinobi.jp': {
    url: 'url',
    ref: 'referrer',
  },

  'googleads.g.doubleclick.net': {
    url: 'url',
    ref: 'ref',
    android: 'msid',
    ios: '_package_name',
    content_url: 'content_url',
  },

  'pubads.g.doubleclick.net': {
    url: 'url',
    ref: 'ref',
    android: 'msid',
    ios: '_package_name',
    content_url: 'content_url',
  },

  'securepubads.g.doubleclick.net': {
    url: 'url',
    ref: 'ref',
    android: 'msid',
    ios: '_package_name',
    content_url: 'content_url',
  },

  'd.socdm.com': {
    url: 'tp',
    ref: 'ref',
    android: 'appbundle',
    ios: 'appbundle',
  },

  'jbbs.shitaraba.net': {
    url: 'url',
  },

  'a.t.webtracker.jp': {
    url: 'url',
  },

  'ssl.webtracker.jp': {
    url: 'url',
  },

  'megalodon.jp': {
    url: 'url',
  },

  'adw.addlv.smt.docomo.ne.jp': {
    url: '_url',
    ref: '_ref',
  },

  's.yimg.jp': {
    url: 'u',
    ref: 'ref',
  },

  'i.yimg.jp': {
    url: 'u',
    ref: 'ref',
  },

  'showads.pubmatic.com': {
    url: 'pageURL',
    ref: 'refurl',
  },

  'optimized-by.rubiconproject.com': {
    url: 'rf',
  },

  'ad.deqwas-dsp.net': {
    url: 'url',
    ref: 'ref',
    domain: 'domain',
  },

  'krad20.deqwas.net': {
    url: 'u',
    domain: 'domain',
  },

  'bidresult-dsp.ad-m.asia': {
    url: 'rf',
  },
};
