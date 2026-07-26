package store

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"strings"
)

func pathPrefixes(p string) map[string]bool {
	result := map[string]bool{constant.Slash: true}
	current := constant.Slash

	for _, segment := range strings.Split(path.Clean(p), constant.Slash) {
		if segment == "" {
			continue
		}

		current = join.Empty(current, segment, constant.Slash)
		result[current] = true
	}

	return result
}
