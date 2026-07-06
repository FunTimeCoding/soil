package maps

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/base64"
)

func ToMarkup(m map[string]string) string {
	result := ""

	for k, v := range m {
		result = fmt.Sprintf("%s  %s: %s\n", result, k, base64.Encode(v))
	}

	return result
}
