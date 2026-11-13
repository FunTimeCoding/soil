package module_graph

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/system"
	"go/build"
	"io"
	"io/fs"
	"os"
)

func forestContext(f *warm_forest.Forest) build.Context {
	result := build.Default
	result.ReadDir = func(directory string) ([]fs.FileInfo, error) {
		if t, relative, found := f.FindTree(directory); found {
			return t.ReadDirectory(relative)
		}

		return diskDirectory(directory)
	}
	result.OpenFile = func(p string) (io.ReadCloser, error) {
		t, relative, found := f.FindTree(p)

		if !found {
			return os.Open(p)
		}

		b, e := t.Read(relative)

		if e != nil {
			return nil, e
		}

		return io.NopCloser(bytes.NewReader(b)), nil
	}
	result.IsDir = func(p string) bool {
		if t, relative, found := f.FindTree(p); found {
			e, known := t.Stat(relative)

			return known && e.Directory
		}

		return system.DirectoryExists(p)
	}

	return result
}
