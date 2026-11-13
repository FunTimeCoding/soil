package scan

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"strings"
)

func ignoredDirectory(name string) bool {
	return name == systemConstant.Temporary ||
		strings.HasPrefix(name, constant.IgnoredPrefix)
}
