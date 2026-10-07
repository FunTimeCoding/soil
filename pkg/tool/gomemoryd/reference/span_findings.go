package reference

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func spanFindings(
	span string,
	d *pointer.Declared,
	r *resolver.Resolver,
	named func(
		scope string,
		name string,
	) bool,
) []*Finding {
	if strings.HasPrefix(span, constant.MemoryScheme) {
		if cited(span, named) {
			return nil
		}

		return []*Finding{NewFinding(span, constant.MissingMemoryText)}
	}

	bare := len(d.Bases) > 0 && pointer.IsBareName(span)

	if !bare && !pointer.IsPath(span) {
		return nil
	}

	var result []*Finding

	for _, c := range pointer.Expand(pointer.NewSpan(span)) {
		if bare {
			if f := underBases(c.Span, d.Bases, r); f != nil {
				result = append(result, f)
			}

			continue
		}

		v := r.Resolve("", d, c)

		if v.Verdict != lintConstant.VerdictLive &&
			v.Verdict != lintConstant.VerdictTallied {
			result = append(
				result,
				NewFinding(c.Span, pointer.Message(v.Verdict, v.Hint)),
			)
		}
	}

	return result
}
