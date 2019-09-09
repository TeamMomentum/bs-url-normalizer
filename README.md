# bs-url-normalizer

bs-url-normalizer は Go の URL 正規化パッケージです。
[Momentum](https://www.m0mentum.co.jp) は自社のサービスで bs-url-normalizer を使用しています。

bs-url-normalizer では Linux と macOS 向けに C 言語から利用できる共有ライブラリを用意しています。 examples ディレクトリにいくつかの言語から bs-url-normalizer を呼び出すサンプルを作成しています。

> bs-url-normalizer normalize a URL by [Moementum](https://www.m0mentum.co.jp). It is used by Momentum products.
>
> You can generate Shared Library of bs-url-normalizer for Linux or macOS. You can see examples for some languages in examples/ directory.


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


## 処理概要

正規化処理として行っているのは以下の処理となります。

* クエリパラメータの順序を統一
    * クエリキーの文字列の値の昇順でソートしております。
    * クエリキーのsliceを引数にsort関数をかけております。
* SPとPCのホスト変換
* 不要なクエリパラメータの除去
* パス末尾の統一
* http/https schemeのhttpへの統一
* パス階層での正規化

> TODO: English

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md)

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (Apache 2.0).
