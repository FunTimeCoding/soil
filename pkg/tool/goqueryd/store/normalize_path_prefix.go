package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func normalizePathPrefix(prefix string) string {
	prefix = strings.TrimLeft(prefix, constant.Slash)

	if prefix != "" && !strings.HasSuffix(prefix, constant.Slash) {
		prefix = fmt.Sprintf("%s/", prefix)
	}

	return prefix
}
