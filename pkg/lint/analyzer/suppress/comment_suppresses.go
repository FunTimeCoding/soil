package suppress

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func CommentSuppresses(
	text string,
	key string,
) bool {
	body, found := strings.CutPrefix(text, constant.CommentPrefix)

	if !found {
		return false
	}

	fields := strings.Fields(body)

	for i, f := range fields {
		if f != constant.IgnoreMarker {
			continue
		}

		if i+1 == len(fields) || !isKey(fields[i+1]) {
			return true
		}

		if fields[i+1] == key {
			return true
		}
	}

	return false
}
