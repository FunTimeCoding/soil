package dash

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func ToUnderscore(s string) string {
	return strings.ReplaceAll(s, separator.Dash, separator.Underscore)
}
