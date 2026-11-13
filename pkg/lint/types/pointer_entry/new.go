package pointer_entry

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_dependency"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

func New(
	root string,
	key string,
	report *file_report.Report,
	spans []*unchecked.Unchecked,
	dependencies *pointer_dependency.Set,
) *Entry {
	return &Entry{
		Root:         root,
		Key:          key,
		Report:       report,
		Unchecked:    spans,
		Dependencies: dependencies,
	}
}
