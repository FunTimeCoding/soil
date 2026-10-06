package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func addPointerConcern(
	s *file_report.Report,
	verdict constant.Verdict,
	hint string,
	path string,
	number int,
	line string,
) {
	s.AddConcern(
		constant.VerdictKeys[verdict],
		pointer.Message(verdict, hint),
		path,
		number,
		line,
		false,
	)
}
