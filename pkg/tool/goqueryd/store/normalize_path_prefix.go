package store

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func normalizePathPrefix(prefix string) string {
	prefix = strings.TrimLeft(prefix, stringsConstant.Slash)

	if prefix != "" && !strings.HasSuffix(prefix, stringsConstant.Slash) {
		prefix = fmt.Sprintf("%s/", prefix)
	}

	return prefix
}
