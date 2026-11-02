package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"strings"
)

func (r *Resolver) ancestor(
	source string,
	normalized string,
) (string, bool) {
	directory, base := path.Split(normalized)
	directory = strings.TrimSuffix(directory, stringsConstant.Slash)
	prefixed := isPrefixShaped(base)
	containing := path.Dir(source)
	candidates := append([]string{containing}, pointer.Ancestors(source)...)

	for i, ancestor := range candidates {
		full := join.Empty(ancestor, stringsConstant.Slash, normalized)

		if i > 0 && r.Exists(full) {
			return full, true
		}

		if !prefixed {
			continue
		}

		parent := ancestor

		if directory != "" {
			parent = join.Empty(ancestor, stringsConstant.Slash, directory)
		}

		if r.Exists(parent) &&
			r.PrefixExists(parent, join.Empty(base, stringsConstant.Dash)) {
			return "", true
		}
	}

	return "", false
}
