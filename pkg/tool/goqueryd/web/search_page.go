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
	status, e := s.service.Status()

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	content := []gomponents.Node{
		html.H3(gomponents.Text(constant.SearchTitle)),
		search_page.Form(query, collection, status.Collections),
	}

	if query != "" {
		results, f := s.searchResults(
			q,
			query,
			collection,
			values.Get(constant.Mode),
		)

		if f != nil {
			s.view.Answer(w, q, s.reporter, f)

			return
		}

		content = append(content, results)
	}

	s.view.RenderPage(
		w,
		constant.SearchTitle,
		webConstant.SearchPath,
		content...,
	)
}
