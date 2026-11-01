package unit

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func goFile(name string) bool {
	return strings.HasSuffix(name, soil.GoExtension)
}
