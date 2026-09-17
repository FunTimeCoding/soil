package front_matter

import (
	"github.com/funtimecoding/soil/pkg/markup/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Extract(content string) (*Front, bool) {
	opening := join.Empty(constant.FrontMatterDelimiter, stringsConstant.Unix)

	if !strings.HasPrefix(content, opening) {
		return nil, false
	}

	rest := content[len(opening):]
	closing := join.Empty(stringsConstant.Unix, constant.FrontMatterDelimiter)
	offset := 0

	for {
		i := strings.Index(rest[offset:], closing)

		if i == -1 {
			return nil, false
		}

		end := offset + i
		after := rest[end+len(closing):]

		if after == "" || strings.HasPrefix(after, stringsConstant.Unix) {
			raw := rest[:end]

			return &Front{
				Raw:   raw,
				Lines: strings.Count(raw, stringsConstant.Unix) + 3,
			}, true
		}

		offset = end + 1
	}
}
