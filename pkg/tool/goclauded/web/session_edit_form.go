package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) sessionEditForm(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	e, f := s.service.GetSession(identifier)
	errors.PanicOnError(f)

	if e == nil {
		s.view.RenderPage(
			w,
			"Session Not Found",
			constant.SessionsPath,
			html.P(gomponents.Text("Session not found.")),
		)

		return
	}

	s.view.RenderPage(
		w,
		"Edit Session",
		constant.SessionsPath,
		html.H3(gomponents.Text("Edit Session")),
		html.Form(
			gomponents.Attr(
				"method",
				"POST",
			),
			gomponents.Attr(
				"action",
				fmt.Sprintf("/sessions/%s/edit", identifier),
			),
			html.Div(
				html.Label(
					gomponents.Attr("for", constant.Alias),
					gomponents.Text("Name"),
				),
				html.Input(
					html.Type("text"),
					html.Name(constant.Alias),
					html.Value(e.AliasValue()),
				),
			),
			html.Div(
				html.Label(
					gomponents.Attr("for", constant.Description),
					gomponents.Text("Description"),
				),
				html.Textarea(
					html.Name(constant.Description),
					html.Rows("3"),
					gomponents.Text(e.Description),
				),
			),
			html.Div(
				html.Button(
					html.Type("submit"),
					gomponents.Text("Save"),
				),
				gomponents.Text(" "),
				html.A(
					gomponents.Attr(
						"href",
						fmt.Sprintf("/sessions/%s", identifier),
					),
					gomponents.Text("Cancel"),
				),
			),
		),
	)
}
