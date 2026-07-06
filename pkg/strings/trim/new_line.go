package trim

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func NewLine(s string) string {
	return strings.TrimRight(s, separator.Unix)
}
