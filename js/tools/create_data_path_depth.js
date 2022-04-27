const fs = require('fs');
const csv = require('csv');
const path = require('path');
const columns = ['host', 'depth', 'query'];

const parser = csv.parse({ columns });
const csvfile = path.join(__dirname, '/../../resources/norm_host_path.csv');
const rs = fs.createReadStream(csvfile, { encoding: 'utf-8' });

console.log('export const pathDepthData = {');
parser.on('readable', () => {
  let data;
  while ((data = parser.read())) {
    const query = data.query;
    if (query === undefined) {
      continue;
    }
    console.log(`  "${data.host}": {`);

    if (query.length > 0) {
      console.log(`    "query": ["${data.query}"],`);
    }

    console.log(`    "depth": ${data.depth}`);
    console.log('  },');
  }
});

parser.on('end', () => {
  console.log('}');
});

rs.pipe(parser);
