import fs from 'fs';
import * as csv from 'csv';
import path from 'path';
import { fileURLToPath } from 'url';
const columns = ['host', 'depth', 'query'];

const parser = csv.parse({ columns });
const dirname = path.dirname(fileURLToPath(import.meta.url));
const csvfile = path.join(dirname, '/../../resources/norm_host_path.csv');
const rs = fs.createReadStream(csvfile, { encoding: 'utf-8' });

console.log(
  `
// @flow
/*:: import type { PathDepth } from './path_depth.js.flow'; */
`.substr(1)
);

console.log('export var N2URLPathDepthData /*: { [string]: PathDepth } */ = {');
parser.on('readable', () => {
  let data;
  while ((data = parser.read())) {
    const query = data.query;
    if (query === undefined) {
      continue;
    }
    console.log(`  '${data.host}': {`);

    if (query.length > 0) {
      console.log(`    query: ['${query}'],`);
    }

    console.log(`    depth: ${data.depth},`);
    console.log('  },');
  }
});

parser.on('end', () => {
  console.log('};');
});

rs.pipe(parser);
