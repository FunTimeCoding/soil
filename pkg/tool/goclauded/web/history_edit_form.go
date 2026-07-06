package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) historyEditForm(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.Atoi(r.PathValue(constant.Identifier))

	if e != nil || identifier <= 0 {
		return
	}

	event := s.service.GetEvent(uint(identifier))

	if event == nil {
		return
	}

	editValue := editableValue(event.Kind, event.Metadata)
	textareaID := fmt.Sprintf("edit-body-%d", identifier)
	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(
		gomponents.Group(
			[]gomponents.Node{
				html.Td(
					gomponents.Text(
						event.CreatedAt.Local().Format("15:04"),
					),
				),
				html.Td(
					html.Textarea(
						html.ID(textareaID),
						html.Name(constant.Body),
						html.Rows("3"),
						html.Style("width: 100%"),
						gomponents.Text(editValue),
					),
					html.Div(
						html.Button(
							gomponents.Attr(
								"hx-post",
								fmt.Sprintf("/history/%d/edit", identifier),
							),
							gomponents.Attr(
								"hx-target",
								fmt.Sprintf("#event-%d", identifier),
							),
							gomponents.Attr("hx-swap", "outerHTML"),
							gomponents.Attr(
								"hx-include",
								fmt.Sprintf("#%s", textareaID),
							),
							gomponents.Text("Save"),
						),
					),
				),
			},
		).Render(w),
	)
}
