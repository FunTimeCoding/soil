package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

func VariableGrouping(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var pending []string
	var pendingStart int

	for s.Scan() {
		line, number := s.Text()

		if isSingleLineVariable(line) {
			if len(pending) == 0 {
				pendingStart = number
			}

			pending = append(pending, line)

			continue
		}

		if len(pending) > 1 {
			flushVariableGroup(s, pending, path, pendingStart)
		} else if len(pending) == 1 {
			s.PassLine(pending[0])
		}

		pending = pending[:0]
		s.PassLine(line)
	}

	if len(pending) > 1 {
		flushVariableGroup(s, pending, path, pendingStart)
	} else if len(pending) == 1 {
		s.PassLine(pending[0])
	}

	return s.Finalize()
}
