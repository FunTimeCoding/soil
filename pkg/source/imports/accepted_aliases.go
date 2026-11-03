package imports

import (
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func acceptedAliases(
	v *import_view.View,
	declaredName string,
	keeper string,
	group []*import_entry.Entry,
) map[*import_entry.Entry]bool {
	result := make(map[*import_entry.Entry]bool)

	if v.Ours == nil {
		return result
	}

	for _, e := range group {
		if e.Path == keeper || e.Local == declaredName || e.Standard {
			continue
		}

		if !v.Ours(e.Module) {
			result[e] = true
		}
	}

	return result
}
