let fs = require('fs');
let csv = require('csv');
let path = require('path');
let columns = ['key', 'value'];

let parser = csv.parse({ columns });
let csvfile = path.join(__dirname, '/../../resources/norm_host_sp.csv');
let rs = fs.createReadStream(csvfile, { encoding: 'utf-8' });

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
