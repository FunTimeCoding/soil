package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Dash(s []string) string {
	return strings.Join(s, constant.Dash)
}
