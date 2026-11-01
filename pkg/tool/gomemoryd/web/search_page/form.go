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
		Inline().
		Search(constant.Query, constant.SearchTitle, query).
		Placeholder(constant.SearchPlaceholder).
		Focus().
		Reveal(
			webConstant.SearchPath,
			join.Empty("#", constant.SearchResultsMark),
		).
		Node()
}
