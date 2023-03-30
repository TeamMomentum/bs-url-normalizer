// @flow
/*:: import type { PathDepth } from './path_depth.js.flow'; */
export var N1URLPathDepthData /*: { [string]: PathDepth } */ = {
  'amigo.gesoten.com': {
    depth: 0,
  },

  'gaingame.gesoten.com': {
    depth: 0,
  },

  'uranai.nosv.org': {
    depth: 1,
    paths: ['/favorite.php'],
  },

  'enjoy.point.auone.jp': {
    depth: 1,
    paths: ['/gacha', '/reward', '/enquete'],
  },

  's.maho.jp': {
    depth: 3,
    paths: ['/book'],
  },

  'ncode.syosetu.com': {
    depth: 2,
    paths: 'ALL',
  },

  'syosetu.org': {
    depth: 3,
    paths: ['/novel'],
  },

  'novel.syosetu.org': {
    depth: 2,
    paths: 'ALL',
  },

  'dokuha.jp': {
    depth: 4,
    paths: ['/comicweb'],
    replace: {
      pattern: '/comicweb/viewer',
      with: '/comicweb/contents',
    },
  },
};
