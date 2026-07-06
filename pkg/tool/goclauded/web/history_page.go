package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) historyPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	kinds := r.URL.Query()[constant.Kind]
	limit := 10
	offset := (page - 1) * limit
	a := argument.Timeline{
		Kinds:  kinds,
		Limit:  limit + 1,
		Offset: offset,
	}
	entries, e := s.service.Timeline(a)
	errors.PanicOnError(e)
	hasMore := len(entries) > limit

	if hasMore {
		entries = entries[:limit]
	}

	table := s.historyTable(entries)
	isHTMX := r.Header.Get("HX-Request") == "true"

	if isHTMX {
		navigation := historyNavigation(page, hasMore, kinds)
		errors.PanicOnError(
			gomponents.Group(
				[]gomponents.Node{table, navigation},
			).Render(w),
		)

		return
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.HistoryTitle)),
		historyFilters(kinds),
	)
	content = append(
		content,
		html.Div(
			html.ID("history-content"),
			table,
			historyNavigation(page, hasMore, kinds),
		),
	)
	s.view.RenderPage(
		w,
		constant.HistoryTitle,
		constant.HistoryPath,
		content...,
	)
}
