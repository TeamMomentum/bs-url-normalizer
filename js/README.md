# URL Normalizer for JavaScript

これは JavaScript 向けの URL 正規化モジュールです。

## インストール

```
$ npm install --save @momentum/url-normalizer
```

## バンドラを用いたビルド

本ライブラリは ES Modules の機構を使って実装されています。rollup などの対応しているバンドラでビルドすることができます。

```
// main.js
import { FirstNormalizedURL } form '@momentum/url-normalizer';

export function OhMyNormalize(s) {
  return FirstNormalizedURL(s);
}
```

```
$ rollup -c --input main.js
```

## ES Modules として読み込む

esm などを使うと Node.js から直接実行できます(Node.jsの `--experimental-modules` には対応していません))

```
// esmodules.js
import { FirstNormalizedURL } form '@momentum/url-normalizer';
const n1url = FirstNormalizedURL('https://www.m0mentum.co.jp/path/to/html');
console.log(n1url);
```

```
$ node -r esm esmodules.js
http://www.m0mentum.co.jp/path/to/html/
```

## CommonJS として読み込む

CommonJS ライブラリとして提供されているので

```
// commonjs.js
const urlnorm = require('@momentum/url-normalizer');
const n1url = urlnorm.FirstNormalizedURL('https://www.m0mentum.co.jp/path/to/html');
console.log(n1url);
```

```
$ node commonjs.js
http://www.m0mentum.co.jp/path/to/html/
```

** examples/js に実際のサンプルがあります **

## BigQuery の UDF として利用する

本ライブラリは BigQuery のUDFとして使用することができます。
Persistent UDF として BigQueryに登録するには以下のようにします。


```
$ make -f bigquery-udf.mk PROJECT=<GCP_PROJECT> DATASET=<BIGQUERY_DATASET> BUCKET_PATH=gs://your-bucket/directory/of/udf
```

登録ができると以下のように利用することができます。

```
$ bq query --nouse_legacy_sql "SELECT <DATASET>.N1URL('https://www.m0mentum.co.jp/path/to/html')"
```
