package match

import (
	"github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/text/literal_edit/block"
)

func region(
	content string,
	b *block.Block,
) (*Match, string) {
	offset, failure := unique(content, b)

	if failure != "" {
		return nil, failure
	}

	switch b.Opening {
	case constant.SpanMarker:
		return span(content, b, offset)
	case constant.AfterMarker, constant.BeforeMarker:
		return insert(content, b, offset), ""
	}

	return &Match{
		Block:   b,
		Offset:  offset,
		Length:  len(b.Search),
		Line:    lineAt(content, offset),
		Replace: b.Replace,
	}, ""
}
