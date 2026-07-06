package status

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func spaces(indent int) string {
	return strings.Repeat(separator.DoubleSpace, indent)
}
