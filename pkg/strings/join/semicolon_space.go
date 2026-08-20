package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func SemicolonSpace(s []string) string {
	return strings.Join(s, constant.SemicolonSpace)
}
