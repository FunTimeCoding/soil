package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/history_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) historyPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	page, e := strconv.Atoi(q.URL.Query().Get(constant.Page))

	if e != nil || page < 1 {
		page = 1
	}

	kinds := q.URL.Query()[constant.Kind]
	entries, f := s.service.Timeline(
		q.Context(),
		argument.Timeline{
			Kinds:  kinds,
			Limit:  constant.HistoryPageSize + 1,
			Offset: (page - 1) * constant.HistoryPageSize,
		},
	)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	more := len(entries) > constant.HistoryPageSize

	if more {
		entries = entries[:constant.HistoryPageSize]
	}

	s.view.RenderPage(
		w,
		constant.HistoryTitle,
		constant.HistoryPath,
		html.H3(gomponents.Text(constant.HistoryTitle)),
		history_page.Filters(kinds),
		history_page.Table(entries, true),
		history_page.Pages(kinds, page, more),
	)
}
