// @flow
/*:: import type { URLInterface } from './url' */
import { parseURL } from './url';
import { adFrameFunc, adFrameParams } from './adframe';

export function createAppURL(
  type /*: 'android' | 'ios' */,
  bundle /*: string */,
  contentUrl /*: ?string */
) /*: URLInterface */ {
  var prefix = type === 'android' ? ':2-' : ':1-';
  var pathname = prefix + bundle;
  var query = {};
  if (contentUrl) {
    query['content_url'] = [encodeURIComponent(contentUrl)];
  }

  return {
    protocol: 'mobileapp:',
    hostname: '',
    port: '',
    pathname: pathname,
    query: query,
    hash: '',
  };
}

// eslint-disable-next-line max-statements, max-lines-per-function
export function getActualPage(url /*: URLInterface */) /*: URLInterface */ {
  if (url.protocol !== 'http:' && url.protocol !== 'https:') {
    return url;
  }

  var query = url.query;
  var f = adFrameFunc[url.hostname];
  if (f) {
    var ret = f(query);
    if (ret) {
      if (ret.type === 'android' || ret.type === 'ios') {
        return createAppURL(ret.type, ret.value);
      }
      url = parseURL(ret.value);
      if (url.protocol !== 'http:' && url.protocol !== 'https:') {
        return url;
      }
    }
  }

  var adframe = adFrameParams[url.hostname];
  if (!adframe) {
    return url;
  }

  var cval = query[adframe.content_url];
  if (cval && cval[0]) {
    var curl = encodeURIComponent(cval[0]);
  }

  var android = query[adframe.android];
  if (android && android[0]) {
    return createAppURL('android', android[0], curl);
  }

  var ios = query[adframe.ios];
  if (ios && ios[0]) {
    return createAppURL('ios', ios[0], curl);
  }

  var uval = query[adframe.url];
  if (uval && uval[0]) {
    url = parseURL(uval[0]);
    return getActualPage(url);
  }

  var domain = query[adframe.domain];
  if (domain && domain[0]) {
    return {
      protocol: url.protocol,
      hostname: domain[0],
      port: '',
      pathname: '',
      query: {},
      hash: '',
    };
  }

  return url;
}

export function getActualRef(url /*: URLInterface */) /*: URLInterface */ {
  var adframe = adFrameParams[url.hostname];
  if (!adframe || !adframe.ref) {
    return url;
  }

  var query = url.query;
  var key = adframe.ref;
  var val = query[key];
  if (val && val[0]) {
    return url;
  }

  return parseURL(val[0]);
}
