package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/block"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func span(
	content string,
	b *block.Block,
	start int,
) (*Match, string) {
	line := lineAt(content, start)
	end := len(content)

	if b.Closing != constant.ToEndMarker {
		from := start + len(b.Search)
		i := strings.Index(content[from:], b.End)

		if i < 0 {
			return nil, fmt.Sprintf(constant.EndNotFound, b.Number, line)
		}

		end = from + i

		if b.Closing == constant.ThroughMarker {
			end += len(b.End)
		}
	}

	replace := b.Replace

	if replace != "" && content[end-1] == '\n' {
		replace = join.Empty(replace, "\n")
	}

	return &Match{
		Block:   b,
		Offset:  start,
		Length:  end - start,
		Line:    line,
		Replace: replace,
	}, ""
}
