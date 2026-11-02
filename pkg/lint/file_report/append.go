package file_report

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func (r *Report) append(
	b *strings.Builder,
	line string,
) {
	b.WriteString(line)
	b.WriteString(stringsConstant.Unix)
}
