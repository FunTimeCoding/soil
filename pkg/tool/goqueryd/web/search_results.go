package web

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_cache"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/search_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchResults(
	q *http.Request,
	page *search_page.Page,
	query string,
	collection string,
	metadata map[string]string,
) (gomponents.Node, error) {
	key := search_cache.Key(query, collection, page.Mode())
	outcome := s.cache.Get(key)

	if outcome == nil {
		o := search_option.New(query, constant.SearchPageSize)
		o.Collection = collection
		o.Mode = page.Mode()
		var e error
		outcome, e = s.service.Search(q.Context(), o)

		if unavailable.Is(e) {
			return html.P(
				html.Class(webConstant.AbsentClass),
				gomponents.Text(e.Error()),
			), nil
		}

		if e != nil {
			return nil, e
		}

		s.cache.Put(key, outcome)
	}

	results, facets := FilterSearchResults(outcome, metadata)

	return page.Results(results, facets), nil
}
