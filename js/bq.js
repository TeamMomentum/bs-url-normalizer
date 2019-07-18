// @flow
/*:: import type { URLInterface } from './url' */

import _URL from 'url-parse';
import { toASCII } from './node_modules/punycode';
import { setURLParser, parseQuery } from './url';

export function parseURL(url /*: string */) /*: URLInterface */ {
  var a = new _URL(url);
  var host = decodeURI(a.hostname);
  var path = decodeURI(a.pathname);
  var hash = decodeURI(a.hash);

  return {
    protocol: a.protocol,
    hostname: toASCII(host),
    port: a.port,
    pathname: encodeURI(path),
    query: parseQuery(a.query),
    hash: encodeURI(hash)
  };
}

setURLParser(parseURL);

export * from './url-normalizer.js';
