package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_cache"
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

	if outcome == nil {
		o := search_option.New(query, 20)
		o.Collection = collection
		o.Mode = mode
		outcome = s.service.Search(o)
		s.cache.Put(key, outcome)
	}

	results, facets := filterSearchResults(outcome, metadata)
	resultsContent := s.renderSearchResults(
		results,
		facets,
		outcome.Degraded,
		query,
		collection,
		metadata,
	)
	isHTMX := r.Header.Get("HX-Request") == "true"

	if isHTMX {
		errors.PanicOnError(
			gomponents.Group(resultsContent).Render(w),
		)

		return
	}

	var content []gomponents.Node
	status := s.service.MustStatus()
	content = append(
		content,
		html.H3(gomponents.Text(constant.SearchTitle)),
		searchForm(query, collection, status.Collections),
		html.Div(
			html.ID("search-results"),
			gomponents.Group(resultsContent),
		),
	)
	s.view.RenderPage(
		w,
		constant.SearchTitle,
		constant.SearchPath,
		content...,
	)
}
