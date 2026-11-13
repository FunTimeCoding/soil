package repository

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Repository) siblingOf(path string) (string, bool) {
	for _, s := range r.Siblings {
		if path == s || strings.HasPrefix(path, join.Empty(s, stringsConstant.Slash)) {
			return s, true
		}
	}

	return "", false
}
