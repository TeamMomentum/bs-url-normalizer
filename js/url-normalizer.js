// @flow
/*:: import type { URLInterface, Query } from './url' */
/*:: import type { PathDepth } from './assets' */

import {
  SPHostData,
  N1URLPathDepthData,
  N2URLPathDepthData,
  IgnoreQueryData,
} from './assets/index.js';
import { createAppURL, getActualPage } from './util.js';
import { parseURL, URLToString } from './url.js';
import { convertUrl } from './convert.js';

// eslint-disable-next-line max-statements
export function FirstNormalizedURL(urlStr /*: string */) /*: string */ {
  var url = createURL(urlStr);
  if (url.protocol !== 'http:') {
    return URLToString(url);
  }

  // it treats triple slash (http:/// or https:///) as empty host and / path
  if (
    (urlStr.indexOf('http:///') === 0 || urlStr.indexOf('https:///') === 0) &&
    url.hostname !== ''
  ) {
    url.pathname = '/' + url.hostname + url.pathname;
    url.hostname = '';
  }

  convertUrl(url);

  var host = url.hostname;
  var m = host.match(/^www\d+\.(.+)/);
  if (m) {
    host = 'www.' + m[1];
  }

  var pathDepth = N1URLPathDepthData[host];
  if (pathDepth) {
    if (pathDepth.replace) {
      url.pathname = url.pathname.replace(
        pathDepth.replace.pattern,
        pathDepth.replace.with
      );
    }

    var path = '';
    if (pathDepth.depth > 0) {
      var cmps = url.pathname.split('/');
      for (var i = 1; i <= pathDepth.depth && i < cmps.length; i++) {
        path += '/' + cmps[i];
      }
    }
    url.pathname = path;
  }

  var len = url.pathname.length;
  if (url.pathname[len - 1] !== '/' && url.protocol !== 'mobileapp:') {
    url.pathname += '/';
  }

  deleteIgnoreQuery(url);

  return URLToString(url);
}

export function SecondNormalizedURL(urlStr /*: string */) /*: string */ {
  var url = secondNormalizedURL(urlStr);
  var str = URLToString(url);
  var len = str.length;
  if (str !== 'http://' && str[len - 1] === '/') {
    return str.substring(0, len - 1);
  }

  return str;
}

function createURL(urlStr /*: string */) /*: URLInterface */ {
  var url = parseURL(urlStr);
  url = getActualPage(url);
  if (url.protocol === 'https:') {
    url.protocol = 'http:';
  }

  var spToPC = SPHostData[url.hostname];
  if (spToPC) {
    url.hostname = spToPC;
  }

  if (url.hostname === 'play.google.com') {
    var query = url.query;
    var v = query.id;
    if (v && v[0]) {
      return createAppURL('android', v[0]);
    }
  }

  if (url.hostname === 'itunes.apple.com') {
    var m = url.pathname.match(/^(\/\w\w)?\/app(\/.+)?\/id(\d+)/);
    if (m) {
      return createAppURL('ios', m[3]);
    }
  }

  url.hostname = trimTailingDots(url.hostname);

  return url;
}

function deleteIgnoreQuery(url /*: URLInterface */) {
  var keys = IgnoreQueryData.ALL.keys;
  var ignore = IgnoreQueryData[url.hostname];
  if (ignore) {
    var paths = ignore.paths;
    var hasPath = paths === 'ALL';
    for (var j = 0; !hasPath && j < paths.length; j++) {
      hasPath = url.pathname.indexOf(paths[j]) >= 0;
    }
    if (hasPath) {
      if (ignore.keys === 'ALL') {
        keys = 'ALL';
      } else {
        // $FlowIgnore
        keys = keys.concat(ignore.keys);
      }
    }
  }

  if (keys === 'ALL') {
    url.query = {};
  } else {
    for (var k in url.query) {
      if (keys.includes(k) || k.indexOf('utm_') === 0) {
        delete url.query[k];
      }
    }
  }
}

function convertToN2URL(
  url /*: URLInterface */,
  path /*: ?string */,
  query /*: ?Query */
) {
  return {
    protocol: url.protocol,
    hostname: url.hostname,
    port: url.port,
    pathname: path || '',
    query: query || {},
    hash: '',
  };
}

// eslint-disable-next-line max-statements
function secondNormalizedURL(urlStr /*: string */) /*: URLInterface */ {
  var url = createURL(urlStr);
  if (url.protocol !== 'http:') {
    return url;
  }

  // it treats triple slash (http:/// or https:///) as empty host
  if (urlStr.indexOf('http:///') === 0 || urlStr.indexOf('https:///') === 0) {
    url.hostname = '';
  }

  convertUrl(url);

  var host = url.hostname;
  var m = host.match(/^www\d+\.(.+)/);
  if (m) {
    host = 'www.' + m[1];
  }

  if (url.protocol === 'mobileapp:') {
    return url;
  }

  var pathDepth = N2URLPathDepthData[host];
  if (pathDepth === undefined) {
    if (url.pathname.indexOf('/~') === 0) {
      pathDepth = ({ depth: 1 } /*: PathDepth */);
    } else {
      pathDepth = ({ depth: 0 } /*: PathDepth */);
    }
  } else if (host === 'www.atwiki.jp') {
    console.log('here', pathDepth, url);
  }

  var path = '';
  if (pathDepth.depth > 0) {
    var cmps = url.pathname.split('/');
    for (var i = 1; i <= pathDepth.depth && i < cmps.length; i++) {
      path += '/' + cmps[i];
    }
  }

  var q /*: { [key: string]: any } */ = {};
  var keys = pathDepth.query;
  if (keys && keys.length > 0) {
    for (var k in url.query) {
      if (keys.includes(k)) {
        q[k] = url.query[k];
      }
    }
  }

  return convertToN2URL(url, path, q);
}

function trimTailingDots(str /*: string */) /*: string */ {
  if (str.endsWith('..')) {
    return str;
  }

  if (str.length > 2 && str[str.length - 1] === '.') {
    return str.substring(0, str.length - 1);
  }

  return str;
}
