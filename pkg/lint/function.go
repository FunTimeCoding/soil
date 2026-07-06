package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
	"strings"
)

func Function(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var functionStart int
	var functionBraceDepth int
	var functionLines []string
	var inFunctionBody bool

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "func ") && strings.Contains(
			line,
			"{",
		) {
			functionStart = number
			functionBraceDepth = strings.Count(line, "{") - strings.Count(
				line,
				"}",
			)
			functionLines = []string{line}
			inFunctionBody = functionBraceDepth > 0
		} else if inFunctionBody {
			functionLines = append(functionLines, line)
			functionBraceDepth += strings.Count(line, "{") - strings.Count(
				line,
				"}",
			)

			if functionBraceDepth == 0 {
				if emptyFunctionBody(functionLines) {
					fixedFunction := toSingleLine(functionLines)
					s.ChangedLine(fixedFunction)
					s.AddConcern(
						constant.EmptyFunctionBodyKey,
						constant.EmptyFunctionBodyText,
						path,
						functionStart,
						join.NewLine(functionLines),
						true,
					)
				} else {
					for _, functionLine := range functionLines {
						s.PassLine(functionLine)
					}
				}

				inFunctionBody = false
				functionLines = nil

				continue
			}
		}

		if !inFunctionBody {
			s.PassLine(line)
		}
	}

	s.Fix = func() {
		if s.Fixed != "" {
			fmt.Printf("Fix functions %s\n", path)
			system.SaveFile(path, s.Fixed)
		}
	}

	return s.Finalize()
}
