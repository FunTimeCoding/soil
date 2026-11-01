package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Region(
	query string,
	results []record.SearchResult,
) gomponents.Node {
	return html.Div(
		html.ID(constant.SearchResultsMark),
		Results(query, results),
	)
}
