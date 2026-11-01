package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/filter"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) entries(
	w http.ResponseWriter,
	r *http.Request,
) {
	q := r.URL.Query()
	f := filter.New()
	f.System = q.Get(constant.System)
	f.Service = q.Get(constant.Service)
	f.User = q.Get(constant.User)

	if t, e := time.Parse(constant.FilterFormat, q.Get(constant.Since)); e == nil {
		f.Since = t
	}

	if t, e := time.Parse(constant.FilterFormat, q.Get(constant.Until)); e == nil {
		f.Until = t
	}

	entries, e := s.service.Entries(f)
	errors.PanicOnError(e)
	table := entry_page.Table(entries)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, table)

		return
	}

	s.view.RenderPage(
		w,
		constant.EntriesTitle,
		constant.EntriesPath,
		html.H1(gomponents.Text(constant.EntriesTitle)),
		entry_page.Filter(q),
		html.Div(html.ID(constant.EntriesMark), table),
	)
}
