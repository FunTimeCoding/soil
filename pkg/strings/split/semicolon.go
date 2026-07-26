package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Semicolon(s string) []string {
	return strings.Split(s, constant.Semicolon)
}
