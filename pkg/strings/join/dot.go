package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Dot(s []string) string {
	return strings.Join(s, separator.Dot)
}
