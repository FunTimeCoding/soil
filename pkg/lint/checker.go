package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

type Checker func(
	string,
	io.Reader,
) *file_report.Report
