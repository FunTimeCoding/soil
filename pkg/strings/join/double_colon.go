package join

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func DoubleColon(s []string) string {
	return strings.Join(s, separator.DoubleColon)
}
