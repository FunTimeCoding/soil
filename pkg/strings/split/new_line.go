package split

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func NewLine(s string) []string {
	return strings.Split(s, separator.Unix)
}
