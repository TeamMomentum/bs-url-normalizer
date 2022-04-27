// @flow
/*:: import type { URLInterface } from './url' */

import _URL from 'url-parse';
import { toASCII } from 'tr46';

export function parseURL(url /*: string */) /*: URLInterface */ {
  var a = new _URL(url);
  var host, path, hash;

  try {
    host = decodeURI(a.hostname);
  } catch (e) {
    host = a.hostname;
  }

  try {
    path = decodeURI(a.pathname);
  } catch (e) {
    path = a.pathname;
  }

  try {
    hash = decodeURI(a.hash);
  } catch (e) {
    hash = a.hash;
  }

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
