package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) entries(
	w http.ResponseWriter,
	r *http.Request,
) {
	f := store.NewFilter()

	if v := r.URL.Query().Get(constant.System); v != "" {
		f.System = v
	}

	if v := r.URL.Query().Get(constant.Service); v != "" {
		f.Service = v
	}

	if v := r.URL.Query().Get(constant.User); v != "" {
		f.User = v
	}

	if v := r.URL.Query().Get(constant.Since); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			f.Since = t
		}
	}

	if v := r.URL.Query().Get(constant.Until); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			f.Until = t
		}
	}

	entries, e := s.store.List(f)
	errors.PanicOnError(e)
	table := entriesTable(entries)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, table)

		return
	}

	s.view.RenderPage(
		w,
		constant.EntriesTitle,
		constant.EntriesPath,
		html.H1(gomponents.Text(constant.EntriesTitle)),
		html.Form(
			html.Class("filter-form"),
			htmx.Get(constant.EntriesPath),
			htmx.Target("#entries-table"),
			htmx.Swap("innerHTML"),
			html.Div(
				html.Class("grid"),
				html.Label(
					gomponents.Text("System"),
					html.Input(
						html.Type("text"),
						html.Name(constant.System),
						html.Value(f.System),
						html.Placeholder("e.g. worker1"),
					),
				),
				html.Label(
					gomponents.Text("Service"),
					html.Input(
						html.Type("text"),
						html.Name(constant.Service),
						html.Value(f.Service),
						html.Placeholder("e.g. nginx"),
					),
				),
				html.Label(
					gomponents.Text("User"),
					html.Input(
						html.Type("text"),
						html.Name(constant.User),
						html.Value(f.User),
					),
				),
			),
			html.Div(
				html.Class("grid"),
				html.Label(
					gomponents.Text("Since"),
					html.Input(
						html.Type("datetime-local"),
						html.Name(constant.Since),
					),
				),
				html.Label(
					gomponents.Text("Until"),
					html.Input(
						html.Type("datetime-local"),
						html.Name(constant.Until),
					),
				),
				html.Label(
					gomponents.Raw("&nbsp;"),
					html.Button(
						html.Type("submit"),
						gomponents.Text("Filter"),
					),
				),
			),
		),
		html.Div(html.ID("entries-table"), table),
	)
}
