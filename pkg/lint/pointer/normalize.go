package pointer

import (
	"fmt"
	"strings"
)

func Normalize(s string) string {
	result := strings.TrimPrefix(s, fmt.Sprintf("%s/", PluginRootVariable))
	result = strings.TrimPrefix(result, "./")

	if i := strings.Index(result, "#"); i != -1 {
		result = result[:i]
	}

	return strings.TrimSuffix(result, "/")
}
