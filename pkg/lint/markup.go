package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	markupConstant "github.com/funtimecoding/soil/pkg/markup/constant"
	"io"
)

func Markup(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)

	for s.Scan() {
		line, number := s.Text()

		if number == 1 {
			if line != markupConstant.FrontMatterDelimiter {
				s.ChangedLine(markupConstant.FrontMatterDelimiter)
				s.ChangedLine(line)
				s.AddConcern(
					constant.FrontMatterDelimiterKey,
					constant.FrontMatterDelimiterText,
					path,
					number,
					line,
					true,
				)
			}
		} else {
			s.PassLine(line)
		}
	}

	return s.Finalize()
}
