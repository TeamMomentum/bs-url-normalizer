## URL Normalizing Module

### How to call normalizing function

You will make by Makefile. It will build with [buildmode=c-shared](https://golang.org/cmd/go/#hdr-Description_of_build_modes) option,
then there will be a shared library "libmomentum\_url\_normalizer.a" on the directory.

Please see the examples on the "examples" directory to see how to use this library.

### Interfaces

There are normalizing functions and util functions as the shared library.
But you can import it directly from "lib/urls", if you are using Go Language


#### First Normalize Function

* Shared

  ```c
  first_normalize_url(char* src, void** dst)
  ```

* Go

  ```go
  func FirstNormalizeURL(*url.URL) string
  ```

#### Second Normalize Function

* Shared

  ```c
  second_normalize_url(char* src, void** dst)
  ```

* Go

  ```go
  func SecondNormalizeURL(*url.URL) string
  ```

#### Deallocating Function

* Shared

  ```c
  free_normalize_url(void* dst)
  ```
* Go

  It will be GC automatically.
