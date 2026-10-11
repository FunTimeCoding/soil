package listing

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/index"
	"path/filepath"
	"strings"
)

func Indexes(root string) ([]*Listing, error) {
	paths, e := filepath.Glob(
		filepath.Join(root, "*", "*", "*", constant.AlpineIndexArchive),
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
		result = append(result, New(parts[0], parts[1], parts[2], entries))
	}

	return result, nil
}
