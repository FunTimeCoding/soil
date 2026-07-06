package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
	"regexp"
	"strings"
)

var (
	errorPattern  = regexp.MustCompile(`\berr(?:\s*(?::=|=[^=])|,)`)
	stringLiteral = regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)
)

func Variable(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "//") {
			continue
		}

		stripped := stringLiteral.ReplaceAllString(trimmed, `""`)

		if errorPattern.MatchString(stripped) {
			s.AddConcern(
				constant.ErrorVariableKey,
				constant.ErrorVariableText,
				path,
				number,
				line,
				false,
			)
		}
	}

	return s.Finalize()
}
