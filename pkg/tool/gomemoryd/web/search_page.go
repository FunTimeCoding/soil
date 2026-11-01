package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/search_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	query := q.URL.Query().Get(constant.Query)
	var results []record.SearchResult

	if query != "" {
		var e error
		results, e = s.service.VisibleSearchResults(
			PrefixQuery(query),
			constant.SearchLimit,
			"",
			"",
			constant.AllScope,
		)

		if e != nil {
			s.view.Answer(w, q, s.reporter, e)

			return
		}
	}

	if s.view.IsExtendedRequest(q) {
		s.view.RenderFragment(w, search_page.Results(query, results))

		return
	}

	s.view.RenderPage(
		w,
		constant.SearchTitle,
		webConstant.SearchPath,
		html.H3(gomponents.Text(constant.SearchTitle)),
		search_page.Form(query),
		search_page.Region(query, results),
	)
}
