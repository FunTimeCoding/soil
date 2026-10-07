package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"github.com/funtimecoding/soil/pkg/alpine/types/listing"
	"path/filepath"
	"strings"
)

func Indexes(root string) ([]*listing.Listing, error) {
	paths, e := filepath.Glob(
		filepath.Join(root, "*", "*", "*", constant.IndexArchive),
	)

	if e != nil {
		return nil, e
	}

	var result []*listing.Listing

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
			listing.New(parts[0], parts[1], parts[2], entries))
	}

	return result, nil
}
