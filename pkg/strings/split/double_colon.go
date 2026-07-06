package split

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func DoubleColon(s string) []string {
	return strings.Split(s, separator.DoubleColon)
}
