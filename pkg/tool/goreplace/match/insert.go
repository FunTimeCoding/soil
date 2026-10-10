package match

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func insert(
	content string,
	b *block.Block,
	anchor int,
) *Match {
	at := strings.LastIndex(content[:anchor], "\n") + 1
	text := join.Empty(b.Replace, "\n")

	if b.Opening == constant.AfterMarker {
		i := strings.Index(content[anchor:], "\n")

		if i < 0 {
			at = len(content)
			text = join.Empty("\n", b.Replace)
		} else {
			at = anchor + i + 1
		}
	}

	return &Match{
		Block:   b,
		Offset:  at,
		Line:    lineAt(content, anchor),
		Replace: text,
		Anchor:  anchor,
	}
}
