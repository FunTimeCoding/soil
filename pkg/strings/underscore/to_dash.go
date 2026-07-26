package underscore

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ToDash(s string) string {
	return strings.ReplaceAll(s, constant.Underscore, constant.Dash)
}
