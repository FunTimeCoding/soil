package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) editSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(r.ParseForm())
	alias := r.FormValue(constant.Alias)
	description := r.FormValue(constant.Description)
	session := s.service.Resolve(identifier)

	if session.Identifier == "" {
		return
	}

	a := edit_session.New()
	a.Alias = &alias
	a.Description = &description
	errors.PanicOnError(s.service.EditSession(session.Identifier, a))
	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	w.Header().Set("HX-Trigger", "session-edited")
	errors.PanicOnError(
		html.Div(
			html.Class("edit-panel"),
			html.P(gomponents.Text(fmt.Sprintf("Saved: %s", alias))),
			html.P(
				html.A(
					gomponents.Attr(
						"hx-get",
						fmt.Sprintf("/conversations/%s", identifier),
					),
					gomponents.Attr("hx-target", "#panel"),
					gomponents.Text("View conversation"),
				),
			),
		).Render(w),
	)
}
