package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) editSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
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
	w.Header().Set(webConstant.ContentType, webConstant.MarkupUnicode)
	w.Header().Set("HX-Trigger", "session-edited")
	errors.PanicOnError(
		gomponents.Group(
			[]gomponents.Node{
				html.P(gomponents.Text(fmt.Sprintf("Saved: %s", alias))),
				load.New(panelPath(identifier), "View conversation").
					Swap("#panel", webConstant.SwapInner).
					Node(),
			},
		).Render(w),
	)
}
