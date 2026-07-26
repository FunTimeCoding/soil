package join

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func CommaSpace(s []string) string {
	return strings.Join(s, constant.CommaSpace)
}
