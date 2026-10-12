package import_view

import "github.com/funtimecoding/soil/pkg/source/types/import_entry"

type View struct {
	PackagePath string
	Entries     []*import_entry.Entry
	LocalTaken  func(string) bool
}
