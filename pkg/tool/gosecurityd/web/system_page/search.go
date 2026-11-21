package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maps"
	"maragu.dev/gomponents"
	"slices"
)

func search(s *system_selection.Selection) gomponents.Node {
	e := entry.New(constant.SystemsPath, "Search").
		Query().
		Inline().
		Search(constant.SearchField, constant.SearchLabel, s.Search).
		Placeholder(constant.SearchLabel)
	kept := s.WithSearch("").Query()

	for _, key := range slices.Sorted(maps.Keys(kept)) {
		for _, value := range kept[key] {
			e.Hidden(key, value)
		}
	}

	return e.Node()
}
