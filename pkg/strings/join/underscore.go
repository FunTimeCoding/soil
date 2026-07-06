package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Underscore(s []string) string {
	return strings.Join(s, separator.Underscore)
}
