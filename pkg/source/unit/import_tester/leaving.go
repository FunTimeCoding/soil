package import_tester

import "github.com/funtimecoding/soil/pkg/source/types/import_entry"

func Leaving(e *import_entry.Entry) *import_entry.Entry {
	e.Stays = false

	return e
}
