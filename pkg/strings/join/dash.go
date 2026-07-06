package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Dash(s []string) string {
	return strings.Join(s, separator.Dash)
}
