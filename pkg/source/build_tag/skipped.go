package build_tag

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func skipped(name string) bool {
	return name == "vendor" ||
		name == "testdata" ||
		name == "tmp" ||
		name != constant.CurrentDirectory &&
			strings.HasPrefix(name, constant.CurrentDirectory)
}
