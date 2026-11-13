package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversation_page"
	"net/http"
)

func (s *Server) conversationEditForm(
	w http.ResponseWriter,
	q *http.Request,
) {
	identifier := q.PathValue(constant.Identifier)
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

	alias := session.Slug
	var description string
	e, found, f := s.roster.FindSession(session.Identifier)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return
	}

	if found {
		if e.Alias != nil {
			alias = *e.Alias
		}

		description = e.Description
	}

	s.view.RenderFragment(
		w,
		conversation_page.EditForm(identifier, alias, description),
	)
}
