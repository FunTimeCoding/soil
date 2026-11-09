package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
)

func Unless(
	skips []string,
	c Checker,
) Checker {
	return func(
		path string,
		r io.Reader,
	) *file_report.Report {
		if Skipped(skips, path) {
			return file_report.New(path, r).Finalize()
		}

		return c(path, r)
	}
}
