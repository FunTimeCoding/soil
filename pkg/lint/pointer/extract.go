package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func Extract(line string) []*Candidate {
	var result []*Candidate
	var prose []string
	parts := strings.Split(line, "`")

	for i, part := range parts {
		if i%2 == 0 || i == len(parts)-1 {
			prose = append(prose, part)

			continue
		}

		if IsPath(part) {
			result = append(result, NewSpan(part))
		}
	}

	for _, m := range constant.LinkTarget.FindAllStringSubmatch(
		strings.Join(prose, "`"),
		-1,
	) {
		if IsPath(m[1]) || IsBareLink(m[1]) || IsAnchor(m[1]) {
			result = append(result, NewLink(unescape(m[1])))
		}
	}

	return result
}
