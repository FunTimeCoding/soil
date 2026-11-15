package system

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func IsGoFile(name string) bool {
	return strings.HasSuffix(name, soil.GoExtension)
}
