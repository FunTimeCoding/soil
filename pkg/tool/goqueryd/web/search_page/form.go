package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func (p *Page) Form(collections []status.CollectionStatus) gomponents.Node {
	choices := []*choice.Choice{choice.New("", "All collections")}

	for _, c := range collections {
		choices = append(choices, choice.New(c.Name, c.Name))
	}

	e := entry.New(webConstant.SearchPath, constant.SearchTitle).
		Query().
		Inline().
		Hidden(constant.Mode, p.mode)

	if p.mode == constant.KeywordMode {
		e.Swap(constant.ResultsTarget, webConstant.SwapInner).Live()
	}

	return e.Search(webConstant.ParameterQuery, constant.SearchTitle, p.query).
		Placeholder("search documents...").
		Focus().
		Select(constant.Collection, "Collection", choices, p.collection).
		Node()
}
