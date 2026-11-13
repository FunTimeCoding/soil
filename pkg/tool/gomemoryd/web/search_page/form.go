package search_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func Form(query string) gomponents.Node {
	return entry.New(webConstant.SearchPath, constant.SearchTitle).
		Query().
		Swap(
			join.Empty("#", constant.SearchResultsMark),
			webConstant.SwapInner,
		).
		Live().
		Inline().
		Search(constant.Query, constant.SearchTitle, query).
		Placeholder(constant.SearchPlaceholder).
		Focus().
		Node()
}
