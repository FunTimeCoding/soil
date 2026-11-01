package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Form(
	query string,
	collection string,
	collections []status.CollectionStatus,
) gomponents.Node {
	choices := []*choice.Choice{choice.New("", "All collections")}

	for _, c := range collections {
		choices = append(choices, choice.New(c.Name, c.Name))
	}

	return entry.New(webConstant.SearchPath, constant.SearchTitle).
		Query().
		Inline().
		Search(webConstant.ParameterQuery, constant.SearchTitle, query).
		Placeholder("search documents...").
		Focus().
		Select(constant.Collection, "Collection", choices, collection).
		Node()
}
