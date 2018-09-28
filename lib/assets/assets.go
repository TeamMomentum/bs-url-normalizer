package assets

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/rakyll/statik/fs"
)

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
// This is for compatibility with go-bindata (https://github.com/jteeuwen/go-bindata)
func Asset(name string) ([]byte, error) {
	stk, err := fs.New() // stl: http.FileSystem
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(name, string(filepath.Separator)) {
		// https://golang.org/pkg/net/http/#FileSystem
		// > A FileSystem implements access to a collection of named files.
		// > The elements in a file path are separated by slash ('/', U+002F) characters, regardless of host operating system convention.
		name = "/" + name
	}
	f, err := stk.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close() // nolint
	return ioutil.ReadAll(f)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// This is for compatibility with go-bindata (https://github.com/jteeuwen/go-bindata)
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic(err)
	}
	return a
}
