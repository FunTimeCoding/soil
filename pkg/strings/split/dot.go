package split

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Dot(s string) []string {
	return strings.Split(s, separator.Dot)
}
