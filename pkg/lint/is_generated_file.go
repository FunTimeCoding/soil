package lint

import "github.com/funtimecoding/soil/pkg/source"

func IsGeneratedFile(path string) bool {
	return source.IsGeneratedFile(path)
}
