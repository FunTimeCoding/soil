package break_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gobreak/edit_span"
	"golang.org/x/tools/cover"
)

func uncovered(
	spans []*edit_span.Span,
	profiles []*cover.Profile,
) *edit_span.Span {
	for _, s := range spans {
		if !s.Uncovered(profiles) {
			return nil
		}
	}

	if len(spans) == 0 {
		return nil
	}

	return spans[0]
}
