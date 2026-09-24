package measure

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func firstLine(content string) string {
	line, _, _ := strings.Cut(content, constant.Unix)

	return line
}
