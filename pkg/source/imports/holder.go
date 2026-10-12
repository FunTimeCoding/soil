package imports

import (
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
	"slices"
)

func holder(
	v *import_view.View,
	group []*import_entry.Entry,
	name string,
) string {
	for _, e := range v.Entries {
		if !e.Stays || slices.Contains(group, e) {
			continue
		}

		if e.Local == name {
			return e.Path
		}
	}

	return ""
}
