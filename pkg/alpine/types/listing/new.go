package listing

import "github.com/funtimecoding/soil/pkg/alpine/index"

func New(
	version string,
	repository string,
	architecture string,
	packages []*index.Entry,
) *Listing {
	return &Listing{
		Version:      version,
		Repository:   repository,
		Architecture: architecture,
		Packages:     packages,
	}
}
