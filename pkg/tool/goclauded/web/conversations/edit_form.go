package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
	e, found, f := s.service.FindSession(session.Identifier)
	errors.PanicOnError(f)

	if found {
		if e.Alias != nil {
			alias = *e.Alias
		}

		description = e.Description
	}

	w.Header().Set(webConstant.ContentType, webConstant.MarkupUnicode)
	errors.PanicOnError(
		gomponents.Group(
			[]gomponents.Node{
				html.H4(gomponents.Text("Edit Session")),
				entry.New(editPath(identifier), "Save").
					Swap("#panel", webConstant.SwapInner).
					Record().
					Text(constant.Alias, "Name", alias).
					Focus().
					TextArea(constant.Description, "Description", description).
					Lines("3").
					Beside(
						load.New(panelPath(identifier), "Cancel").
							Swap("#panel", webConstant.SwapInner).
							Button().
							Class("outline secondary").
							Node(),
					).
					Node(),
			},
		).Render(w),
	)
}
