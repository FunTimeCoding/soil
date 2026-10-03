package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	"io"
	"slices"
	"strings"
)

func InterruptingList(
	path string,
	r io.Reader,
) *file_report.Report {
	b, e := io.ReadAll(r)
	errors.PanicOnError(e)
	content := string(b)
	s := file_report.New(path, strings.NewReader(content))
	interruptions := reflow.Interruptions(content)

	for s.Scan() {
		line, number := s.Text()
		s.PassLine(line)

		if !slices.Contains(interruptions, number) {
			continue
		}

		s.AddConcern(
			constant.InterruptingListKey,
			constant.InterruptingListText,
			path,
			number,
			line,
			false,
		)
	}

	return s.Finalize()
}
