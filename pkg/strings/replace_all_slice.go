package strings

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ReplaceAllSlice(
	content string,
	replaces []string,
) string {
	for k, v := range ToMap(replaces, constant.Equals) {
		content = strings.ReplaceAll(content, k, v)
	}

	return content
}
