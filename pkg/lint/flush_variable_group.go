package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func flushVariableGroup(
	s *file_report.Report,
	lines []string,
	path string,
	startLine int,
) {
	s.ChangedLine("var (")

	for _, line := range lines {
		s.ChangedLine(join.Empty("\t", strings.TrimPrefix(line, "var ")))
	}

	s.ChangedLine(")")
	s.AddConcern(
		constant.VariableGroupingKey,
		constant.VariableGroupingText,
		path,
		startLine,
		lines[0],
		true,
	)
}
