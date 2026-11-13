package gobreak

import (
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/edit_span"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/overlay"
)

func apply(
	root string,
	modulePath string,
	b *break_file.Break,
	o *overlay.Overlay,
) ([]*edit_span.Span, string) {
	var result []*edit_span.Span

	for _, file := range files(b) {
		spans, failure := applyFile(root, modulePath, file, b, o)

		if failure != "" {
			return nil, failure
		}

		result = append(result, spans...)
	}

	return result, ""
}
