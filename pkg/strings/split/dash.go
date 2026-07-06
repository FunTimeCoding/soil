package split

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Dash(s string) []string {
	return strings.Split(s, separator.Dash)
}
