package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/collection_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) documentsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	name := q.PathValue("name")
	sourceType := q.URL.Query().Get(constant.SourceType)
	page, e := strconv.Atoi(q.URL.Query().Get(constant.Page))

	if e != nil || page < 1 {
		page = 1
	}

	var metadata map[string]string

	if sourceType != "" {
		metadata = map[string]string{constant.SourceType: sourceType}
	}

	outcome, f := s.service.ListDocuments(
		name,
		metadata,
		constant.DocumentPageSize+1,
		(page-1)*constant.DocumentPageSize,
		false,
	)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	results := outcome.Results
	more := len(results) > constant.DocumentPageSize

	if more {
		results = results[:constant.DocumentPageSize]
	}

	s.view.RenderPage(
		w,
		name,
		constant.CollectionsPath,
		html.H3(gomponents.Text(name)),
		collection_page.SourceTypes(name, outcome.Facets, sourceType),
		collection_page.Documents(results),
		collection_page.Pages(name, sourceType, page, more),
	)
}
