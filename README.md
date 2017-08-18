Read this in English [here](https://github.com/TeamMomentum/bs-url-normalizer/blob/master/README.en.md)

## URL正規化モジュール

### 呼び出し方法

Makefileでは[buildmode=c-shared](https://golang.org/cmd/go/#hdr-Description_of_build_modes)を指定しており、makeするとlibmomentum\_url\_normalizer.aというShared Libraryが生成されます。

簡単なサンプルをexamplesディレクトリに用意してありますので、そちらを参考にURL正規化を行ってください。

### インタフェース

正規化関数とメモリ開放関数が実装されており、以下のようなインターフェースとなっております。

なお Go用のライブラリは ./lib/urls 以下に配置しております。

#### 1段階目評価時関数

* Shared

  ```c
  first_normalize_url(char* src, void** dst)
  ```

* Go

  ```go
  func FirstNormalizeURL(*url.URL) string
  ```

#### 2段階目評価時関数

* Shared

  ```c
  second_normalize_url(char* src, void** dst)
  ```

* Go

  ```go
  func SecondNormalizeURL(*url.URL) string
  ```

#### リソース開放関数

* Shared

  ```c
  free_normalize_url(void* dst)
  ```
* Go

  リソースはGCされるためインターフェースを用意しておりません。

### 処理概要

正規化処理として行っているのは以下の処理となります。

* クエリパラメータの順序を統一

  クエリキーの文字列の値の昇順でソートしております。

  クエリキーのsliceを引数にsort関数をかけております。

  ```go
  // import sort
  sort.Strings([]string)
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
