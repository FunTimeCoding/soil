package unit

import "github.com/funtimecoding/soil/pkg/source/types/import_entry"

func leaving(e *import_entry.Entry) *import_entry.Entry {
	e.Stays = false

	return e
}
