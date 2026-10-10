package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func unique(
	content string,
	b *block.Block,
) (int, string) {
	offsets := occurrences(content, b.Search)

	switch len(offsets) {
	case 1:
		return offsets[0], ""
	case 0:
		return 0, notFound(content, b)
	}

	var lines []string

	for _, o := range offsets {
		lines = append(lines, fmt.Sprint(lineAt(content, o)))
	}

	return 0, fmt.Sprintf(
		constant.Ambiguous,
		b.Number,
		len(offsets),
		strings.Join(lines, constant.LineSeparator),
	)
}
