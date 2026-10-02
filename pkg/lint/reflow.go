package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
	"slices"
	"strings"
	"unicode/utf8"
)

func Reflow(
	path string,
	r io.Reader,
) *file_report.Report {
	b, e := io.ReadAll(r)
	errors.PanicOnError(e)
	content := string(b)
	s := file_report.New(path, strings.NewReader(content))
	rewrapped, f := reflow.Reflow(content, constant.ReflowWidth)

	if f != nil {
		s.Add(
			concern.NewFile(
				constant.ReflowRefusedKey,
				join.Space(constant.ReflowRefusedText, f.Error()),
				path,
				false,
			),
		)
	}

	lines := strings.Split(rewrapped, "\n")

	for s.Scan() {
		line, number := s.Text()
		s.PassLine(line)

		if f != nil || utf8.RuneCountInString(line) <= constant.ReflowWidth ||
			slices.Contains(lines, line) {
			continue
		}

		s.AddConcern(
			constant.ReflowKey,
			constant.ReflowText,
			path,
			number,
			line,
			true,
		)
	}

	result := s.Finalize()

	if f == nil && rewrapped != content {
		result.Fixed = rewrapped
	}

	return result
}
