package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func Colon(s ...string) string {
	return strings.Join(s, separator.Colon)
}
