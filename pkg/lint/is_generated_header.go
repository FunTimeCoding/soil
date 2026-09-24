package lint

import "github.com/funtimecoding/soil/pkg/source"

func IsGeneratedHeader(content string) bool {
	return source.IsGeneratedHeader(content)
}
