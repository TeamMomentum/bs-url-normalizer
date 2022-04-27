const fs = require('fs');
const csv = require('csv');
const path = require('path');
const columns = ['key', 'value'];

const parser = csv.parse({ columns });
const csvfile = path.join(__dirname, '/../../resources/norm_host_sp.csv');
const rs = fs.createReadStream(csvfile, { encoding: 'utf-8' });

console.log('export const spHostData = {');
parser.on('readable', () => {
  let data;
  while ((data = parser.read())) {
    // eslint-disable-line no-cond-assign
    console.log(`  "${data.key}": "${data.value}",`);
  }
});

parser.on('end', () => {
  console.log('}');
});

rs.pipe(parser);
