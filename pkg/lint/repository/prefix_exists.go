package repository

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Repository) PrefixExists(
	directory string,
	prefix string,
) bool {
	full := join.Empty(directory, constant.Slash, prefix)

	for _, p := range r.Files.Files() {
		if strings.HasPrefix(p, full) {
			return true
		}
	}

	return false
}
