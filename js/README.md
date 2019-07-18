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

## 開発

```
js/
├── Makefile
├── README.md
├── adframe/            # adframe を正規化するためのjsファイル群
├── assets/             # 特定のURLを正規化するためのjsファイル群
├── bigquery-udf.mk     # BigQuery UDF を作るための Makefile
├── bq.js               # BigQuery UDF を作るためのjs
├── build/
├── package-lock.json
├── package.json
├── prettier.config.js
├── rollup.config.js
├── test/
├── tools/              # assets/ ディレクトリのファイルを作成するためのjsスクリプト
├── url-normalizer.js   # メインファイル。正規化の処理はここ
├── url.js
└── util.js
```

### assets/params.js

ホスト毎に以下のようなプロパティを持つオブジェクトを返します。値はクエリパラメータのキーです。

- url: "実際に広告が表示されたURL",
- ref: "実際に広告が表示されたURLのReferrer",
- android: "アンドロイドのパッケージネーム",
- ios: "iOSのパッケージネーム",
- content_url: "android または ios の場合に、リンクする Web 上のURL"

