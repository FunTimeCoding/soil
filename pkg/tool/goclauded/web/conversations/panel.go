package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) panel(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	session := s.service.Resolve(identifier)

	if session.Identifier == "" {
		w.Header().Set(web.ContentType, "text/html; charset=utf-8")
		errors.PanicOnError(
			html.P(gomponents.Text("Session not found.")).Render(w),
		)

		return
	}

	messages := s.service.Messages(session.Identifier)
	var nodes []gomponents.Node

	for i := range messages {
		if messages[i].IsMeta {
			continue
		}

		nodes = append(nodes, messageBlock(&messages[i]))
	}

	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(gomponents.Group(nodes).Render(w))
}
