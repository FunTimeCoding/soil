package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/argument/edit_session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) conversationEditSubmit(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
	alias := q.FormValue(constant.Alias)
	description := q.FormValue(constant.Description)
	session := s.transcript.Resolve(identifier)

	if session.Identifier == "" {
		s.view.Answer(
			w,
			q,
			s.reporter,
			not_found.New(constant.SessionKind, identifier),
		)

		return
	}

	a := edit_session.New()
	a.Alias = &alias
	a.Description = &description

	if e := s.roster.EditSession(session.Identifier, a); e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	w.Header().Set(webConstant.TriggerHeader, constant.SessionEditedEvent)
	s.view.RenderFragment(w, conversation_page.Saved(identifier, alias))
}
