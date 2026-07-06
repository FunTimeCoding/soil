package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func CommaSpace(s []string) string {
	return strings.Join(s, separator.CommaSpace)
}
