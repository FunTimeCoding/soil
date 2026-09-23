package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func Normalize(s string) string {
	result := strings.TrimPrefix(s, constant.PluginRootPrefix)
	result = strings.TrimPrefix(result, "./")

	if i := strings.Index(result, "#"); i != -1 {
		result = result[:i]
	}

	result = constant.LineSuffix.ReplaceAllString(result, "")

	return strings.TrimSuffix(result, "/")
}
