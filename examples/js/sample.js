const { BSURLNormalizer } = require('../../js/bs-url-normalizer.js');

let aurl = BSURLNormalizer.URLParse("https://www.google.com");
console.log(aurl);
console.log(BSURLNormalizer.FirstNormalizedURL(aurl));
