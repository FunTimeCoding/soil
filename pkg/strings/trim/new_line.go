package trim

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func NewLine(s string) string {
	return strings.TrimRight(s, constant.Unix)
}
