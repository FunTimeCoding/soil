package build_tag

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func skipped(name string) bool {
	return name == "vendor" ||
		name == "testdata" ||
		name == "tmp" ||
		name != soil.CurrentDirectory &&
			strings.HasPrefix(name, soil.CurrentDirectory)
}
