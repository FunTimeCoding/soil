package dash

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ToUnderscore(s string) string {
	return strings.ReplaceAll(s, constant.Dash, constant.Underscore)
}
