package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Slash(s []string) string {
	return strings.Join(s, separator.Slash)
}
