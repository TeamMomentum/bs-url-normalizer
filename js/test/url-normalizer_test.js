// @flow
/*::
type Test = {
  in: string;
  n1url?: string;
  n2url?: string;
  r1url?: string;
  r2url?: string;
};
*/

import { strict as assert } from 'assert';
import fs from 'fs';
import path from 'path';

// $FlowIgnore
import { before, describe, it } from 'mocha';
import { FirstNormalizedURL, SecondNormalizedURL } from '../url-normalizer.js';

// $FlowIgnore
const dirname = path.dirname(new URL(import.meta.url).pathname);

describe('normalizer', () => {
  const dir = path.join(dirname, '../../testdata');
  const ignores = process.argv
    .filter((a) => a.startsWith('--ignore='))
    .map((a) => a.substring(9).split(','))
    .flat();
  const files = fs.readdirSync(dir).filter((f) => !ignores.includes(f));

  console.log('ignores', ignores);

  it('has some file', () => assert.ok(files.length > 0));
  files.forEach((f) =>
    describe(f, () => {
      let tests /*: Array<Test> */ = [];
      before(() => {
        const data = fs.readFileSync(path.join(dir, f), { encoding: 'utf-8' });
        tests = JSON.parse(data).tests;
      });

      it('has some test', () => assert.ok(tests.length > 0));
      it('tests', () => testNorm(tests));
    })
  );
});

function testNorm(tests /*: Array<Test> */) {
  tests.forEach((t) =>
    describe(t.in, () => {
      it('needs a input', () => assert(t.in));
      it(
        'is valid test',
        () => assert.ok(t.n1url || t.n2url || t.r1url || t.r2url),
        JSON.stringify(t)
      );

      if (t.n1url) {
        it('n1url', () => assert.equal(t.n1url, FirstNormalizedURL(t.in)));
      }

      if (t.n2url) {
        it('n2url', () => assert.equal(t.n2url, SecondNormalizedURL(t.in)));
      }

      if (t.r1url) {
        it('r1url is not implemented yet');
      }

      if (t.r2url) {
        it('r2url is not implemented yet');
      }
    })
  );
}
