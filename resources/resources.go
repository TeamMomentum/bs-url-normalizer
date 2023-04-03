package resources

import (
	"embed"
)

//go:embed *
var resources embed.FS

func MustAsset(name string) []byte {
	data, err := resources.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return data
}
