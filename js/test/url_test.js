// @flow
import _assert from 'assert';

import { describe, it } from 'mocha';
import { parseURL } from '../url';

const assert = (_assert /*: any */).strict;

describe('url.parseURL', () => {
  const q = '&empty&empty2=&q=query&クエリ=日本語&%E3%82%AF%E3%82%A8%E3%83%AA=%E6%97%A5%E6%9C%AC';
  const s = 'http://日本語.com:8080/path/日本語/%E6%97%A5%E6%9C%AC?' + q;
  const src = s + '&url=' + encodeURIComponent(s) + '#' + q;

  console.log('src', src);

  const want = {
    protocol: 'http:',
    hostname: 'xn--wgv71a119e.com',
    port: '8080',
    pathname: '/path/%E6%97%A5%E6%9C%AC%E8%AA%9E/%E6%97%A5%E6%9C%AC',
    query: {
      'empty2': [''],
      'q': ['query'],
      'クエリ': ['日本語', '日本'],
      'url': ['http://日本語.com:8080/path/日本語/%E6%97%A5%E6%9C%AC?&empty&empty2=&q=query&クエリ=日本語&%E3%82%AF%E3%82%A8%E3%83%AA=%E6%97%A5%E6%9C%AC']
    },
    hash:
      '#&empty&empty2=&q=query&%E3%82%AF%E3%82%A8%E3%83%AA=%E6%97%A5%E6%9C%AC%E8%AA%9E&%E3%82%AF%E3%82%A8%E3%83%AA=%E6%97%A5%E6%9C%AC'
  };
  const got = (parseURL(src) /*: any */);

  for (const key in want) {
    it(`has valid ${key}`, () => assert.deepEqual(want[key], got[key]));
  }
});
