package warm_tree

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func (t *Tree) Read(p string) ([]byte, error) {
	clean := path.Clean(p)
	e, found := t.entries[clean]

	if !found {
		return nil, fs.ErrNotExist
	}

	if e.Directory {
		return nil, fs.ErrInvalid
	}

	if cached, read := t.content[clean]; read {
		return cached, nil
	}

	b, f := os.ReadFile(filepath.Join(t.root, filepath.FromSlash(clean)))

	if f != nil {
		return nil, f
	}

	t.content[clean] = b

	return b, nil
}
