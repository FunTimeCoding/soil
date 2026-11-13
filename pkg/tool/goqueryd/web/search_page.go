package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	values := q.URL.Query()
	query := values.Get(webConstant.ParameterQuery)
	collection := values.Get(constant.Collection)
	metadata := parseMetadataParams(values)
	page := search_page.New(
		query,
		collection,
		values.Get(constant.Mode),
		metadata,
	)
	var results gomponents.Node

	if query != "" {
		var e error
		results, e = s.searchResults(q, page, query, collection, metadata)

		if e != nil {
			s.view.Answer(w, q, s.reporter, e)

			return
		}
	}

	if s.view.IsExtendedRequest(q) {
		s.view.RenderFragment(w, gomponents.Group([]gomponents.Node{results}))

		return
	}

	status, f := s.service.Status()

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	s.view.RenderPage(
		w,
		constant.SearchTitle,
		webConstant.SearchPath,
		html.H3(gomponents.Text(constant.SearchTitle)),
		page.Form(status.Collections),
		page.Modes(),
		html.Div(html.ID(constant.ResultsIdentifier), results),
	)
}
