var urlRegexpSrc = "^(?<schema>http(s)?://)?"
  + "((?<user>[^:]+):?(?<password>.+)?@)?"
  + "(?<host>[^#?:/]+)(:(?<port>[0-9]+))?"
  + "(?<path1>/[^#?/]+)?"
  + "(?<path2>/[^#?/]+)?"
  + "(?<path3>/[^#?/]+)?"
  + "(?<path4>/[^#?/]+)?"
  + "(?<path5>/[^#?/]+)?"
  + "[^?#]*"
  + "(?<query>\\?[^#]+)?"
  + "(?<hash>#.+)?$"
var urlRegExp = new RegExp(urlRegexpSrc);
function SecondNormalizedURL(url) {
  return 'http://' + url.match(urlRegExp).groups.host;
}

var m = "http://username:password@www.google.co.jp:8080/path/to/html?q=hi+hey+hoo#title".match(urlRegExp)
console.log(m)
