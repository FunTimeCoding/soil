package warm_tree

import (
	"github.com/funtimecoding/soil/pkg/source/types/named_entry"
	"io/fs"
	"path"
)

func (t *Tree) ReadDirectory(p string) ([]fs.FileInfo, error) {
	clean := path.Clean(p)
	e, found := t.entries[clean]

	if !found {
		return nil, fs.ErrNotExist
	}

	if !e.Directory {
		return nil, fs.ErrInvalid
	}

	var result []fs.FileInfo

	for _, name := range t.children[clean] {
		result = append(
			result,
			named_entry.New(name, t.entries[path.Join(clean, name)]),
		)
	}

	return result, nil
}
