package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func Normalize(s string) string {
	result := strings.TrimPrefix(s, constant.PluginRootPrefix)
	result = strings.TrimPrefix(result, "./")
	result = constant.LineSuffix.ReplaceAllString(result, "")

	return strings.TrimSuffix(result, "/")
}
