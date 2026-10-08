package segment

import (
	"github.com/funtimecoding/soil/pkg/lint/types/segment_span"
	"strings"
)

func segmentSpans(name string) []*segment_span.Span {
	var result []*segment_span.Span
	offset := 0

	for partIndex, part := range strings.Split(name, "_") {
		if partIndex > 0 {
			offset++
		}

		r := []rune(part)
		start := 0

		for i := 1; i < len(r); i++ {
			if !boundary(r, i) {
				continue
			}

			result = append(
				result,
				segment_span.New(
					offset+len(string(r[:start])),
					offset+len(string(r[:i])),
					strings.ToLower(string(r[start:i])),
				),
			)
			start = i
		}

		if start < len(r) {
			result = append(
				result,
				segment_span.New(
					offset+len(string(r[:start])),
					offset+len(part),
					strings.ToLower(string(r[start:])),
				),
			)
		}

		offset += len(part)
	}

	return result
}
