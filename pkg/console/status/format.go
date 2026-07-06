package status

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"slices"
	"strings"
)

func (s *Status) Format() string {
	result := key_value.Empty(
		spaces(s.format.Indentation),
		strings.Join(s.bubbles, " | "),
	)

	for _, t := range s.linesByTag {
		if slices.Contains(s.format.Tags, t.Tag) {
			result = extendLines(result, t.Line, s.format.Indentation)
		}
	}

	if s.format.ShowExtended && len(s.lines) > 0 {
		result = extendLines(result, s.lines, s.format.Indentation)
	}

	return result
}
