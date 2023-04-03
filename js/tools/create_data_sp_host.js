import fs from 'fs';
import * as csv from 'csv';
import path from 'path';
import { fileURLToPath } from 'url';
const columns = ['key', 'value'];

const parser = csv.parse({ columns });
const dirname = path.dirname(fileURLToPath(import.meta.url));
const csvfile = path.join(dirname, '/../../resources/norm_host_sp.csv');
const rs = fs.createReadStream(csvfile, { encoding: 'utf-8' });

console.log('export const SPHostData = {');
parser.on('readable', () => {
  let data;
  while ((data = parser.read())) {
    // eslint-disable-line no-cond-assign
    console.log(`  '${data.key}': '${data.value}',`);
  }
});

parser.on('end', () => {
  console.log('};');
});

rs.pipe(parser);
