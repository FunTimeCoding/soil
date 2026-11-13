package checker_step

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

type Step struct {
	Name  string
	Check func(string, io.Reader) *file_report.Report
	Pure  bool
}
