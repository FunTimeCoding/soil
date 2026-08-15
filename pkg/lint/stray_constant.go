package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
	"path/filepath"
	"strings"
)

func StrayConstant(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	base := filepath.Base(path)

	if base == "constant.go" ||
		constantDirectory(path) {
		for s.Scan() {
			line, _ := s.Text()
			s.PassLine(line)
		}

		return s.Finalize()
	}

	var packageConstant bool
	var sanction bool
	var depth int

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(line, "package ") {
			packageConstant = strings.TrimPrefix(line, "package ") == "constant"
		}

		if !packageConstant && !sanction && depth == 0 &&
			(strings.HasPrefix(trimmed, "const ") || trimmed == "const (") {
			s.AddConcern(
				constant.StrayConstantKey,
				constant.StrayConstantText,
				path,
				number,
				line,
				false,
			)
		}

		rule, okay := parseFixtureRule(line)
		sanction = okay &&
			rule == constant.StrayConstantRule &&
			testdataPath(path)

		for _, c := range line {
			switch c {
			case '{':
				depth++
			case '}':
				depth--
			}
		}

		s.PassLine(line)
	}

	return s.Finalize()
}
