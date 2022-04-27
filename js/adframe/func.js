// @flow
// refer ../../lib/urls/optimize.go

/*::
import type { Query } from '../url'
type Type = 'android' | 'ios' | 'web'
type Result = {
  type: Type,
  value: string
}

*/

export var adFrameFunc /*: { [string]: (Query) => ?Result } */ = {
  'd.socdm.com': function (query) {
    var tp = query.sdktype;
    if (typeof tp === 'undefined' || typeof tp[0] === 'undefined') {
      return null;
    }

    switch (tp[0]) {
      case '0':
        var v0 = query.url;
        if (!v0 || !v0[0]) {
          return null;
        }
        return {
          type: 'web',
          value: v0[0],
        };
      case '1':
        var v1 = query.appbundle;
        if (!v1 || !v1[0]) {
          return null;
        }
        return {
          type: 'android',
          value: v1[0],
        };
      case '2':
        var v2 = query.appbundle;
        if (!v2 || !v2[0]) {
          return null;
        }
        return {
          type: 'ios',
          value: v2[0],
        };
    }

    return null;
  },
};
