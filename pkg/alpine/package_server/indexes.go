package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"path/filepath"
	"strings"
)

func Indexes(root string) ([]*Listing, error) {
	paths, e := filepath.Glob(
		filepath.Join(root, "*", "*", "*", constant.IndexArchive),
	)

	if e != nil {
		return nil, e
	}

	var result []*Listing

	for _, path := range paths {
		entries, f := index.Read(path)

		if f != nil {
			return nil, f
		}

		relative, g := filepath.Rel(root, path)

		if g != nil {
			return nil, g
		}

		parts := strings.Split(relative, string(filepath.Separator))
		result = append(
			result,
			&Listing{
				Version:      parts[0],
				Repository:   parts[1],
				Architecture: parts[2],
				Packages:     entries,
			},
		)
	}

	return result, nil
}
