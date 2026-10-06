package heading

import (
	"github.com/funtimecoding/soil/pkg/markup/front_matter"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func body(content string) string {
	f, found := front_matter.Extract(content)

	if !found {
		return content
	}

	lines := strings.SplitAfter(content, constant.Unix)

	return strings.Join(lines[f.Lines:], "")
}
