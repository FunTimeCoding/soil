package installed

import (
	"debug/buildinfo"
	"github.com/funtimecoding/soil/pkg/stamp"
	"path/filepath"
)

func Read(path string) (*Binary, bool) {
	i, e := buildinfo.ReadFile(path)

	if e != nil {
		return nil, false
	}

	s := stamp.Read(i)
	result := New(filepath.Base(path), path)
	result.Module = s.Module
	result.Version = s.Version

	return result, true
}
