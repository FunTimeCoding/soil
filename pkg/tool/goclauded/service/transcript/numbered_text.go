package transcript

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func numberedText(read string) string {
	lines := strings.Split(read, stringsConstant.Unix)

	for i := range lines {
		lines[i] = constant.NumberedLine.ReplaceAllString(lines[i], "")
	}

	return join.NewLine(lines)
}
