package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

func Check(
	path string,
	r io.Reader,
) *file_report.Report {
	s := New(path, r)

	for s.report.Scan() {
		line, number := s.report.Text()
		s.step(line, number)
	}

	return s.report.Finalize()
}
