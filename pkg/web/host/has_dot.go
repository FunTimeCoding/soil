package host

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func HasDot(s string) bool {
	return strings.Count(s, separator.Dot) > 0
}
