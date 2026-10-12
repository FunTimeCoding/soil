package service

import "github.com/funtimecoding/soil/pkg/source/types/import_entry"

func heldByOther(
	entries []*import_entry.Entry,
	arriving *import_entry.Entry,
	name string,
) bool {
	for _, e := range entries {
		if e != arriving && e.Stays && e.Local == name {
			return true
		}
	}

	return false
}
