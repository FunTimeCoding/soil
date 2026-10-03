package reference

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func Check(
	content string,
	r *pointer.Resolver,
	named func(
		scope string,
		name string,
	) bool,
) []*Finding {
	var result []*Finding
	declared := pointer.NewDeclared()

	for _, line := range strings.Split(content, stringsConstant.Unix) {
		parts := strings.Split(line, string(lintConstant.Backtick))

		for i, part := range parts {
			if i%2 == 0 || i == len(parts)-1 {
				result = append(result, barePaths(part, r.Roots)...)

				continue
			}

			if strings.HasPrefix(part, constant.MemoryScheme) {
				if !cited(part, named) {
					result = append(
						result,
						NewFinding(part, constant.MissingMemoryText),
					)
				}

				continue
			}

			if !pointer.IsPath(part) {
				continue
			}

			for _, c := range pointer.Expand(pointer.NewSpan(part)) {
				v := r.Resolve("", declared, c)

				if v.Verdict == lintConstant.VerdictLive ||
					v.Verdict == lintConstant.VerdictTallied {
					continue
				}

				result = append(
					result,
					NewFinding(c.Span, lintConstant.VerdictTexts[v.Verdict]),
				)
			}
		}
	}

	return result
}
