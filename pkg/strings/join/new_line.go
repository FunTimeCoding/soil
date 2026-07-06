package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func NewLine(s []string) string {
	return strings.Join(s, separator.Unix)
}
