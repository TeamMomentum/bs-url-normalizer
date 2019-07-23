// @flow
/*::
export interface URLInterface {
  protocol: string;
  hostname: string;
  port: string;
  pathname: string;
  query: Query;
  hash: string;
}

export type Query = { [string]: Array<string> }
*/

var URLParser = DefaultURLParser;

export function DefaultURLParser(s /*: string */) /*: URLInterface */ {
  var a;
  if (typeof URL !== 'undefined') {
    a = new URL(s);
  } else if (
    typeof window !== 'undefined' &&
    typeof window.document !== 'undefined'
  ) {
    a = window.document.createElement('a');
  } else {
    throw new Error('cannot parse URL');
  }

  return {
    protocol: a.protocol,
    hostname: a.hostname,
    port: a.port,
    pathname: a.pathname,
    query: parseQuery(a.search),
    hash: a.hash
  };
}

export function setURLParser(f /*: (string) => URLInterface */) {
  URLParser = f;
}

export function parseURL(s /*: string */) /*: URLInterface */ {
  return URLParser(s);
}

export function URLToString(url /*: URLInterface */) /*: string */ {
  if (!url) {
    return '';
  }

  var scheme = url.protocol;
  switch (scheme) {
    case 'mobileapp:':
      return url.protocol + url.pathname;
    case 'http:':
    case 'https:':
    case 'ftp:':
    case 'file:':
      scheme += '//';
  }

  var host = url.hostname;
  if (url.port) {
    host += ':' + url.port;
  }

  return (
    scheme + host + url.pathname + orderedQueryString(url.query) + url.hash
  );
}

function orderedQueryString(query /*: Query */) /*: string */ {
  if (!query) {
    return '';
  }

  var keys = [];
  for (var k in query) {
    keys.push(k);
  }
  keys = keys.sort();
  if (keys.length === 0) {
    return '';
  }

  var arr = [];
  for (var i = 0; i < keys.length; i++) {
    var key = keys[i];
    var vals = query[key].sort();
    for (var j = 0; j < vals.length; j++) {
      var p = encodeURIComponent(key) + '=' + encodeURIComponent(vals[j]);
      arr.push(p);
    }
  }
  return '?' + arr.join('&');
}

export function parseQuery(queryString /*: string */) /*: Query */ {
  var query = {};
  if (typeof queryString !== 'string' || queryString.length === 0) {
    return query;
  }
  var params = queryString.substr(1).split('&');
  for (var i = 0; i < params.length; i++) {
    if (!params[i]) {
      // skip: ?&cb=0.9359973386788167
      continue;
    }
    var kv = params[i].split('=');
    if (kv.length !== 2) {
      continue;
    }

    var k = kv[0];
    // 文字コードの問題で URIError が発生する可能性があるので、ダメな時は諦める
    try {
      k = decodeURIComponent(k);
    } catch (e) {}

    var v = kv[1];
    // 文字コードの問題で URIError が発生する可能性があるので、ダメな時は諦める
    try {
      v = decodeURIComponent(v);
    } catch (e) {}

    if (typeof query[k] === 'undefined') {
      query[k] = [];
    }
    query[k].push(v);
  }

  return query;
}
