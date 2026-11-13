package module_graph

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func excludedName(name string) bool {
	return name == "testdata" ||
		name == "vendor" ||
		strings.HasPrefix(name, soil.CurrentDirectory) ||
		strings.HasPrefix(name, "_")
}
