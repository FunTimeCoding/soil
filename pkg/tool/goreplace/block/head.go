package block

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"slices"
)

func (b *Block) head(
	lines []string,
	start int,
) (int, error) {
	if b.Opening != constant.SpanMarker {
		search, next, marker := collect(lines, start, constant.DividerMarker)

		if marker == "" {
			return 0, fmt.Errorf(constant.MissingDivider, b.Number)
		}

		b.Search = search

		return next, nil
	}

	search, next, marker := collect(
		lines,
		start,
		append(slices.Clone(constant.Closings), constant.DividerMarker)...,
	)

	if marker == "" || marker == constant.DividerMarker {
		return 0, fmt.Errorf(constant.MissingClosing, b.Number)
	}

	b.Search = search
	b.Closing = marker
	end, after, divider := collect(lines, next, constant.DividerMarker)

	if divider == "" {
		return 0, fmt.Errorf(constant.MissingDivider, b.Number)
	}

	b.End = end

	return after, nil
}
