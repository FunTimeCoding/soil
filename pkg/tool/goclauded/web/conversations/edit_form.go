package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) editForm(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	session := s.service.Resolve(identifier)
	alias := session.Slug
	var description string
	e, f := s.service.GetSession(session.Identifier)
	errors.PanicOnError(f)

	if e != nil {
		if e.Alias != nil {
			alias = *e.Alias
		}

		description = e.Description
	}

	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(
		html.Div(
			html.Class("edit-panel"),
			html.H4(gomponents.Text("Edit Session")),
			html.Form(
				gomponents.Attr(
					"hx-post",
					fmt.Sprintf("/conversations/%s/edit", identifier),
				),
				gomponents.Attr("hx-target", "#panel"),
				html.Div(
					html.Class("edit-field"),
					html.Label(
						gomponents.Attr("for", constant.Alias),
						gomponents.Text("Name"),
					),
					html.Input(
						html.Type("text"),
						html.Name(constant.Alias),
						html.Value(alias),
						gomponents.Attr("autofocus", ""),
					),
				),
				html.Div(
					html.Class("edit-field"),
					html.Label(
						gomponents.Attr("for", constant.Description),
						gomponents.Text("Description"),
					),
					html.Textarea(
						html.Name(constant.Description),
						html.Rows("3"),
						gomponents.Text(description),
					),
				),
				html.Div(
					html.Class("edit-actions"),
					html.Button(
						html.Type("submit"),
						gomponents.Text("Save"),
					),
					html.Button(
						html.Type("button"),
						gomponents.Attr(
							"hx-get",
							fmt.Sprintf(
								"/conversations/%s",
								identifier,
							),
						),
						gomponents.Attr("hx-target", "#panel"),
						gomponents.Text("Cancel"),
					),
				),
			),
		).Render(w),
	)
}
