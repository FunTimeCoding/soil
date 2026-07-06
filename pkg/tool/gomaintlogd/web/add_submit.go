package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) addSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	e := entry.New()
	e.Action = r.FormValue(constant.Action)
	e.User = r.FormValue(constant.User)
	e.System = r.FormValue(constant.System)
	e.Service = r.FormValue(constant.Service)
	e.Description = r.FormValue(constant.Description)
	errors.PanicOnError(s.store.Add(e))
	s.view.RenderFragment(
		w,
		html.Div(
			html.Div(
				html.Class("success"),
				gomponents.Textf(
					"Entry added: %s by %s",
					e.Action,
					e.User,
				),
			),
			addForm(),
		),
	)
}
