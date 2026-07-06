package split

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Semicolon(s string) []string {
	return strings.Split(s, separator.Semicolon)
}
