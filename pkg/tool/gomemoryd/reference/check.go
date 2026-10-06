package reference

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func Check(
	content string,
	bases []string,
	r *pointer.Resolver,
	named func(
		scope string,
		name string,
	) bool,
) []*Finding {
	declared, result := declare(bases, r)

	for _, line := range strings.Split(content, stringsConstant.Unix) {
		parts := strings.Split(line, string(lintConstant.Backtick))

		for i, part := range parts {
			if i%2 == 0 || i == len(parts)-1 {
				result = append(result, barePaths(part, r.Roots)...)

				continue
			}

			result = append(result, spanFindings(part, declared, r, named)...)
		}
	}

	return result
}
