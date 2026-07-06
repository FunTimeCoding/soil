package text

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/text/option"
	"strings"
)

func OptimizeWhitespace(
	input string,
	s *option.Whitespace,
) string {
	if s == nil {
		s = option.New()
	}

	var result []string
	var blankLines int

	for _, line := range split.NewLine(input) {
		line = strings.TrimSpace(line)

		if line == "" {
			blankLines++

			if blankLines > s.AllowedBlankLines {
				continue
			}
		} else {
			blankLines = 0
		}

		result = append(result, line)
	}

	if s.NewlineAtEnd {
		// If no newline at the end, add one
		if len(result) > 0 && result[len(result)-1] != "" {
			result = append(result, "")
		}
	}

	return join.NewLine(result)
}
