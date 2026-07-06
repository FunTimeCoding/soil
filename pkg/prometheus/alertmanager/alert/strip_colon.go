package alert

import (
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"strings"
)

func StripColon(s string) string {
	if strings.ContainsRune(s, ':') {
		result, _ := key_value.Colon(s)

		return result
	}

	return s
}
