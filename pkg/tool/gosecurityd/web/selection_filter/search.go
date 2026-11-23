package selection_filter

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maps"
	"maragu.dev/gomponents"
	"slices"
)

func (f *Filter) search() gomponents.Node {
	e := entry.New(f.path, "Search").
		Query().
		Inline().
		Search(constant.SearchField, constant.SearchLabel, f.selection.Search).
		Placeholder(constant.SearchLabel)
	kept := f.query(f.selection.WithSearch(""))

	for _, key := range slices.Sorted(maps.Keys(kept)) {
		for _, value := range kept[key] {
			e.Hidden(key, value)
		}
	}

	return e.Node()
}
