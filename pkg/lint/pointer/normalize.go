package pointer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func Normalize(s string) string {
	result := strings.TrimPrefix(
		s,
		fmt.Sprintf("%s/", constant.PluginRootVariable),
	)
	result = strings.TrimPrefix(result, "./")

	if i := strings.Index(result, "#"); i != -1 {
		result = result[:i]
	}

	return strings.TrimSuffix(result, "/")
}
