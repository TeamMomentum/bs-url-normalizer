const urlnorm = require('@momentum/url-normalizer');
const { FirstNormalizedURL, SecondNormalizedURL } = urlnorm;

const s = "http://example.com/path";
const n1url = FirstNormalizedURL(s);
const n2url = SecondNormalizedURL(s);
console.log("                URL:", s);
console.log(" FirstNormalizedURL:", n1url);
console.log("SecondNormalizedURL:", n2url);
