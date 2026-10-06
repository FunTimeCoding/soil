package build_tag

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"maps"
	"path/filepath"
)

func Files(directory string) map[string][]string {
	if directory == "" {
		directory = constant.CurrentDirectory
	}

	if !memoizing.Load() {
		return scan(directory)
	}

	key, e := filepath.Abs(directory)

	if e != nil {
		key = directory
	}

	if cached, okay := discovered.Load(key); okay {
		return maps.Clone(cached.(map[string][]string))
	}

	result := scan(directory)
	discovered.Store(key, result)

	return maps.Clone(result)
}
