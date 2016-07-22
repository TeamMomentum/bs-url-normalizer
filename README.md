## URL正規化モジュール

### インタフェース

./lib/urls 以下に配置しております。

モジュールはGo言語で書かれており、正規化関数は以下のインタフェースとなっております。

* 1段階目評価時関数
  ```go
  func FirstNormalizeURL(*url.URL) string
  ```

* 2段階目評価時関数
  ```go
  func SecondNormalizeURL(*url.URL) string
  ```

### 処理概要

正規化処理として行っているのは以下の処理となります。

* クエリパラメータの順序を統一

  クエリキーの文字列の値の昇順でソートしております。

  クエリキーのsliceを引数にsort関数をかけております。

  ```go
  // import sort
  sort.Strings([]string)
  ```

  perlのコードで記述した場合、以下と同等となることかと思います。

  ```perl
  use utf8;

  sort @keys;
  ```

* SPとPCのホスト変換

  ```go
  func normalizeSPHost(*url.URL)
  ```

* 不要なクエリパラメータの除去

  ```go
  func removeQueryParameters(*url.URL, url.Values)
  ```

* パス末尾の統一

  ```go
  func normalizePathSuffix(*url.URL)
  ```

* http/https schemeのhttpへの統一

  ```go
  func normalizeScheme(*url.URL)
  ```

* パス階層での正規化

  ```go
  func normalizePath(ul *url.URL) bool
  ```

## テスト

  ```
  $ go test ./lib/... -v
  === RUN   TestRemoveQueryParameters
  --- PASS: TestRemoveQueryParameters (0.00s)
  === RUN   TestQueryOrder
  --- PASS: TestQueryOrder (0.00s)
  === RUN   TestNormalizeURLFormat
  --- PASS: TestNormalizeURLFormat (0.00s)
  === RUN   TestSplitNDomainPath
  --- PASS: TestSplitNDomainPath (0.00s)
  === RUN   TestNormalizePathMap
  --- PASS: TestNormalizePathMap (0.00s)
  PASS
  ok      github.com/TeamMomentum/bs-url-normalizer/lib/urls     0.012s
  ```