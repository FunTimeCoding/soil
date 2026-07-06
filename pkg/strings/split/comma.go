package split

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Comma(s string) []string {
	return strings.Split(s, separator.Comma)
}
