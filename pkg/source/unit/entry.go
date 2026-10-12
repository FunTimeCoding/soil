package unit

import "github.com/funtimecoding/soil/pkg/source/types/import_entry"

func entry(
	path string,
	declared string,
	local string,
) *import_entry.Entry {
	return import_entry.New(path, declared, local, "example.org/alfa", false)
}
