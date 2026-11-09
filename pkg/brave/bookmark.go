package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/bookmark"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"os"
	"path/filepath"
)

func Bookmark(profile string) (*bookmark.File, error) {
	p, e := Profile(profile)

	if e != nil {
		return nil, e
	}

	b, f := os.ReadFile(filepath.Join(p.Path, constant.BookmarksFile))

	if f != nil {
		return nil, f
	}

	return bookmark.Parse(b)
}
