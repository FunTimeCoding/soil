package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"strings"
)

func (r *Resolver) ancestorExists(
	source string,
	normalized string,
) bool {
	directory, base := path.Split(normalized)
	directory = strings.TrimSuffix(directory, constant.Slash)
	prefixed := isPrefixShaped(base)
	containing := path.Dir(source)
	candidates := append([]string{containing}, Ancestors(source)...)

	for i, ancestor := range candidates {
		full := join.Empty(ancestor, constant.Slash, normalized)

		if i > 0 && r.Exists(full) {
			return true
		}

		if !prefixed {
			continue
		}

		parent := ancestor

		if directory != "" {
			parent = join.Empty(ancestor, constant.Slash, directory)
		}

		if r.Exists(parent) &&
			r.PrefixExists(parent, join.Empty(base, constant.Dash)) {
			return true
		}
	}

	return false
}
