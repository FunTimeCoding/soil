package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"strings"
)

func resolveInclude(
	owner string,
	initFile bool,
	item string,
) string {
	if !strings.HasPrefix(item, constant.StateSeparator) {
		return item
	}

	parts := strings.Split(owner, constant.StateSeparator)

	if !initFile {
		parts = parts[:len(parts)-1]
	}

	rest := strings.TrimLeft(item, constant.StateSeparator)
	up := len(item) - len(rest) - 1

	if up > len(parts) {
		up = len(parts)
	}

	parts = parts[:len(parts)-up]

	return strings.Join(append(parts, rest), constant.StateSeparator)
}
