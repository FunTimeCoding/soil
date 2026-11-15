package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func addPointerConcern(
	s *file_report.Report,
	key string,
	hint string,
	path string,
	number int,
	line string,
) {
	s.AddConcern(key, pointer.Message(key, hint), path, number, line, false)
}
