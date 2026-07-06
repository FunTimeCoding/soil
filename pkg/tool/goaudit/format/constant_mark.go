package format

import "github.com/funtimecoding/soil/pkg/tool/goaudit/scan"

func constantMark(s *scan.Service) string {
	if s.ConstantDirectory {
		return "dir"
	}

	if s.ConstantFile {
		return "file"
	}

	return "-"
}
