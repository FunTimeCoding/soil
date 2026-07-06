package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Pipe(s []string) string {
	return strings.Join(s, separator.Pipe)
}
