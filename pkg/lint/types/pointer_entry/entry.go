package pointer_entry

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_dependency"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

type Entry struct {
	Root         string
	Key          string
	Report       *file_report.Report
	Unchecked    []*unchecked.Unchecked
	Dependencies *pointer_dependency.Set
}
