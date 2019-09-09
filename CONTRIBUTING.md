## Development

### Requirements

- [GNU Make](https://www.gnu.org/software/make/)
- [Go](https://golang.org)
- [statick](https://github.com/rakyll/statik): To embed asset files into Go codes

### Building Shared Library

```sh
# `make build` will do:
# 1. update Go dependencies,
# 2. update asset files
# 3. run tests
# 4. build a shared library file
$ make build
```


### Update dependencies

```sh
make dep
```

### Test

```sh
$ make test
go test -v -race ./...
...
```

### Updating embedded asset files (Optional)

```sh
$ make assets
```
