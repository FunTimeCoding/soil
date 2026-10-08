package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func notFound(
	content string,
	b *block.Block,
) string {
	result := fmt.Sprintf(constant.NotFound, b.Number)

	if strings.Contains(collapse(content), collapse(b.Search)) {
		return fmt.Sprintf(constant.HintFormat, result, constant.WhitespaceOnly)
	}

	first := strings.SplitN(b.Search, "\n", 2)[0]
	var lines []string

	for i, line := range strings.Split(content, "\n") {
		if strings.TrimSpace(first) != "" && strings.Contains(line, first) {
			lines = append(lines, fmt.Sprint(i+1))
		}

		if len(lines) == constant.HintLimit {
			break
		}
	}

	if len(lines) == 0 {
		return fmt.Sprintf(
			constant.HintFormat,
			result,
			constant.FirstLineAbsent,
		)
	}

	return fmt.Sprintf(
		constant.HintFormat,
		result,
		fmt.Sprintf(
			constant.FirstLineAt,
			strings.Join(lines, constant.LineSeparator),
		),
	)
}
