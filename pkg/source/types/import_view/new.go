package import_view

import "github.com/funtimecoding/soil/pkg/source/types/import_entry"

func New(
	packagePath string,
	entries []*import_entry.Entry,
	localTaken func(string) bool,
) *View {
	return &View{
		PackagePath: packagePath,
		Entries:     entries,
		LocalTaken:  localTaken,
	}
}
