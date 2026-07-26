package split

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Dash(s string) []string {
	return strings.Split(s, constant.Dash)
}
