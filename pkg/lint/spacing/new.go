package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

func New(
	path string,
	r io.Reader,
) *Spacing {
	return &Spacing{report: file_report.New(path, r), path: path}
}
