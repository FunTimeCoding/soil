package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Comma(s []string) string {
	return strings.Join(s, separator.Comma)
}
