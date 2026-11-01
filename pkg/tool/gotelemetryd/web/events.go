package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web/event_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) events(
	w http.ResponseWriter,
	r *http.Request,
) {
	q := r.URL.Query()
	page := 1

	if n, e := strconv.Atoi(q.Get(constant.PageParameter)); e == nil && n > 0 {
		page = n
	}

	o := query_option.New()
	o.Tool = q.Get(constant.Tool)
	o.Surface = q.Get(constant.Surface)
	o.Actor = q.Get(constant.Actor)
	o.Kind = q.Get(constant.Kind)
	o.Limit = constant.EventPageSize + 1
	o.Offset = (page - 1) * constant.EventPageSize
	events, e := s.service.Query(o)
	errors.PanicOnError(e)
	more := len(events) > constant.EventPageSize

	if more {
		events = events[:constant.EventPageSize]
	}

	s.view.RenderPage(
		w,
		"",
		constant.EventsPath,
		html.H3(gomponents.Text(constant.EventsTitle)),
		event_page.Filters(o),
		event_page.Table(events),
		event_page.Pager(o, page, more),
	)
}
