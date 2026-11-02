package store

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"strings"
)

func pathPrefixes(p string) map[string]bool {
	result := map[string]bool{stringsConstant.Slash: true}
	current := stringsConstant.Slash

	for _, segment := range strings.Split(path.Clean(p), stringsConstant.Slash) {
		if segment == "" {
			continue
		}

		current = join.Empty(current, segment, stringsConstant.Slash)
		result[current] = true
	}

	return result
}
