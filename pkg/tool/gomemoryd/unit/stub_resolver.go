package unit

import "github.com/funtimecoding/soil/pkg/lint/pointer"

func stubResolver() *pointer.Resolver {
	result := pointer.New()
	result.Roots = []string{"doc", "pkg"}
	result.Exists = func(p string) bool {
		return p == "doc/real.md"
	}

	return result
}
