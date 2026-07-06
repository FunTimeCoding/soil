package status

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status/tag_line"
)

func (s *Status) TagLine(
	tag string,
	format string,
	a ...any,
) *Status {
	for _, t := range s.linesByTag {
		if t.Tag == tag {
			t.Line = append(t.Line, fmt.Sprintf(format, a...))

			return s
		}
	}

	t := tag_line.New(tag)
	t.Line = append(t.Line, fmt.Sprintf(format, a...))
	s.linesByTag = append(s.linesByTag, t)

	return s
}
