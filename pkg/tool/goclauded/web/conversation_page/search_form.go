package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/entry/field"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"slices"
)

func searchForm(
	query string,
	kinds []string,
) gomponents.Node {
	e := entry.New(searchPath(), "Search").
		Query().
		Swap(constant.SidebarEntriesTarget, webConstant.SwapInner).
		Live().
		Class("search-form").
		Search(constant.QueryField, "Search", query).
		Placeholder("Search conversations...").
		Group("Kinds")

	for _, kind := range constant.BlockKinds {
		e.Field(
			field.New(
				webConstant.EntryCheckboxKind,
				constant.KindField,
				kind,
				kind,
			).Checked(slices.Contains(kinds, kind)),
		)
	}

	return e.Node()
}
