package measure

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func firstLine(content string) string {
	line, _, _ := strings.Cut(content, stringsConstant.Unix)

	return line
}
