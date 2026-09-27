package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Middot(s []string) string {
	return strings.Join(s, constant.SpacedMiddot)
}
