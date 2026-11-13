package checker_step

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

func New(
	name string,
	check func(string, io.Reader) *file_report.Report,
	pure bool,
) *Step {
	return &Step{Name: name, Check: check, Pure: pure}
}
