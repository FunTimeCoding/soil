package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"path"
	"strings"
)

func pathPrefixes(p string) map[string]bool {
	result := map[string]bool{separator.Slash: true}
	current := separator.Slash

	for _, segment := range strings.Split(path.Clean(p), separator.Slash) {
		if segment == "" {
			continue
		}

		current = join.Empty(current, segment, separator.Slash)
		result[current] = true
	}

	return result
}
