# bs-url-normalizer

bs-url-normalizer は Go の URL 正規化パッケージです。
[Momentum](https://www.m0mentum.co.jp) は自社のサービスで bs-url-normalizer を使用しています。

bs-url-normalizer では Linux と macOS 向けに C 言語から利用できる共有ライブラリを用意しています。 examples ディレクトリにいくつかの言語から bs-url-normalizer を呼び出すサンプルを作成しています。

> bs-url-normalizer normalize a URL by [Moementum](https://www.m0mentum.co.jp). It is used by Momentum products.
>
> You can generate Shared Library of bs-url-normalizer for Linux or macOS. You can see examples for some languages in examples/ directory.

## 処理概要

**NOTE**: 本パッケージによる正規化はデータベースのキーなどに使うことを想定しています。正規化後のURLはブラウザ等から **アクセスできない可能性があります** 。ご留意ください。

正規化には第一正規化と第二正規化があります。

> TODO: English

### 第一正規化

単一のWebページを表すキーとして利用できます。
正規化処理として行っているのは以下の処理となります。

* http/https schemeのhttpへの統一
* SPとPCのホスト変換 [リスト参照](./resources/norm_host_sp.csv)
* パス階層での正規化
* パス末尾を `/` で統一
* 不要なクエリパラメータの除去
* クエリパラメータの順序を統一
    * クエリキーの文字列の値の昇順でソートしております。
    * クエリキーのsliceを引数にsort関数をかけております。
* フラグメントを削除

### 第二正規化

単一のWebサイトを表すキーとして利用できます (bakusai.com など CGM サイト等では例外があります)。
正規化処理として行っているのは以下の処理となります。

* http/https schemeのhttpへの統一
* SPとPCのホスト変換 [リスト参照](./resources/norm_host_sp.csv)
* パス末尾の `/` は削除
* 指定ドメイン、指定パターン以外のパスは全て削除
    * ユーザスペースは保持
    * [リスト参照](./resources/norm_host_path.csv)
* クエリパラメータの除去
    * 特定ドメインにおいては一部を保持
* フラグメントを削除


### 具体例

[testdata](./testdata) にサンプルがありますので参照してください。

- `in`: 正規化前のURLを表します
- `n1url`: 第一正規化を表します
- `n2url`: 第二正規化を表します


## Usage

### Go

#### Functions

```go
func FirstNormalizeURL(*url.URL) string
func SecondNormalizeURL(*url.URL) string
```

#### Example

https://play.golang.org/p/SvgzdnbV6BM

```go
package main

import (
	"net/url"
	"github.com/TeamMomentum/bs-url-normalizer/lib/urls"
)

func main() {
	u, _ := url.Parse("https://www.m0mentum.co.jp/ja/about.html")
	n1url := urls.SecondNormalizeURL(u)
	println(n1url) //=> http://www.m0mentum.co.jp
}
```


### C

#### Functions

```c
extern void first_normalize_url(char* src, void** result);
extern void second_normalize_url(char* src, void** result);
extern void free_normalize_url(void** result);
```

#### Example

関数を呼び出した後は確保したメモリを解放する必要があります

> You have to free a memory that is allocated by bs-url-normalizer

```c
#include <stdio.h>

#include "libmomentum_url_normalizer.h"

int main() {
	void *result;

	second_normalize_url("https://www.m0mentum.co.jp/about.html", &result);
	printf("%s\n", (char*)result) //=> http://www.m0mentum.co.jp;
	free_normalize_url(result);

        return 0;
}
```


## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md)

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (Apache 2.0).
