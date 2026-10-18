package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_cache"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query().Get("query")
	collection := r.URL.Query().Get(constant.Collection)
	mode := r.URL.Query().Get(constant.Mode)
	metadata := parseMetadataParams(r.URL.Query())

	if mode == "" {
		mode = "hybrid"
	}

	if query == "" {
		s.renderSearchForm(w, collection)

		return
	}

	key := search_cache.Key(query, collection, mode)
	outcome := s.cache.Get(key)
	var e error

	if outcome == nil {
		o := search_option.New(query, 20)
		o.Collection = collection
		o.Mode = mode
		outcome, e = s.service.Search(r.Context(), o)

		if e == nil {
			s.cache.Put(key, outcome)
		}
	}

	var resultsContent []gomponents.Node

	if unavailable.Is(e) {
		resultsContent = []gomponents.Node{
			html.P(html.Em(gomponents.Text(e.Error()))),
		}
	} else {
		errors.PanicOnError(e)
		results, facets := FilterSearchResults(outcome, metadata)
		resultsContent = s.renderSearchResults(
			results,
			facets,
			query,
			collection,
			metadata,
		)
	}

	isHTMX := r.Header.Get("HX-Request") == "true"

	if isHTMX {
		errors.PanicOnError(gomponents.Group(resultsContent).Render(w))

		return
	}

	var content []gomponents.Node
	status := s.service.MustStatus()
	content = append(
		content,
		html.H3(gomponents.Text(constant.SearchTitle)),
		searchForm(query, collection, status.Collections),
		html.Div(html.ID("search-results"), gomponents.Group(resultsContent)),
	)
	s.view.RenderPage(
		w,
		constant.SearchTitle,
		webConstant.SearchPath,
		content...)
}
