package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"io"
	"strings"
)

func PackageName(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)

	for s.Scan() {
		line, number := s.Text()
		trimmed := strings.TrimSpace(line)

		if !strings.HasPrefix(trimmed, "package ") {
			continue
		}

		name := strings.Fields(trimmed)[1]

		for _, blocked := range constant.PackageBlocklist {
			if name == blocked {
				s.AddConcern(
					constant.PackageNameKey,
					constant.PackageNameText,
					path,
					number,
					line,
					false,
				)
			}
		}
	}

	return s.Finalize()
}
