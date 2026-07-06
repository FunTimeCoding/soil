package compact

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"strings"
)

func Mail(s string) string {
	if strings.Contains(s, separator.At) {
		first, _ := key_value.At(s)

		return first
	}

	return s
}
