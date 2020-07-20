// @flow
/*:: import type { URLInterface } from './url' */

// convert5chItest convert URL to where the page will be redirected by JavaScript
// e.g "http://itest.5ch.net/foo/test/read.cgi/abc/" => "http://foo.5ch.net/test/read.cgi/abc/"
export function convertUrl(url /*: URLInterface */) {
  convert5chItest(url)
  || convertAppleApps(url);
}

function convert5chItest(url /*: URLInterface */) {
  if (url.hostname !== 'itest.5ch.net' && url.hostname !== 'itest.bbspink.com') {
    return false;
  }

  var m = url.pathname.match("^/([^/]+)(/test/read.cgi/.*)$");
  if (!m) {
    return false;
  }

  url.hostname = url.hostname.replace('itest', m[1]);
  url.pathname = m[2];
  return true;
}

function convertAppleApps(url /*: URLInterface */) {
  if (url.hostname !== 'apps.apple.com') {
    return false;
  }

  var m = url.pathname.match("^.*/app/.*id(\\d+)");
  if (!m) {
    return false;
  }

  url.protocol = 'mobileapp:';
  url.pathname = ':1-' + m[1];
  url.query = {};

  return true;
}
