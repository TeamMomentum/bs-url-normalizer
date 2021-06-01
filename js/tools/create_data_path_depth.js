let fs = require('fs');
let csv = require('csv');
let path = require('path');
let columns = ['host', 'depth', 'query'];

let parser = csv.parse({ columns });
let csvfile = path.join(__dirname, '/../../resources/norm_host_path.csv');
let rs = fs.createReadStream(csvfile, { encoding: 'utf-8' });

console.log('export const pathDepthData = {');
parser.on('readable', () => {
  let data;
  while ((data = parser.read())) {
    // eslint-disable-line no-cond-assign
    let query = data.query;
    if (query === undefined) {
      continue;
    }
    console.log(`  "${data.host}": {`);

    if (query.length > 0) {
      console.log(`    "query": ["${data.query}"],`);
    }

    console.log(`    "depth": ${data.depth}`);
    console.log(`  },`);
  }
});

parser.on('end', () => {
  console.log('}');
});

rs.pipe(parser);
