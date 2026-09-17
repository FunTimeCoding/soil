package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
)

func addPointerConcern(
	s *file_report.Report,
	verdict constant.Verdict,
	path string,
	number int,
	line string,
) {
	s.AddConcern(
		constant.VerdictKeys[verdict],
		constant.VerdictTexts[verdict],
		path,
		number,
		line,
		false,
	)
}
