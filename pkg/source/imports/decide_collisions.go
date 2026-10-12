package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
	"maps"
	"slices"
)

func decideCollisions(v *import_view.View) []*import_decision.Decision {
	byName := make(map[string][]*import_entry.Entry)

	for _, e := range v.Entries {
		if e.Stays && e.Declared != constant.ConstantPackage {
			byName[e.Declared] = append(byName[e.Declared], e)
		}
	}

	var result []*import_decision.Decision

	for _, name := range slices.Sorted(maps.Keys(byName)) {
		if group := byName[name]; len(group) > 1 {
			result = append(result, decideCollision(v, name, group)...)
		}
	}

	return result
}
