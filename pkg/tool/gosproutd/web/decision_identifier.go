package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"net/http"
	"strconv"
)

func (s *Server) decisionIdentifier(
	w http.ResponseWriter,
	q *http.Request,
) (*decision.Decision, bool) {
	raw := q.URL.Query().Get(constant.IdentifierParameter)
	identifier, e := strconv.ParseUint(raw, 10, 64)

	if e != nil {
		s.view.Answer(
			w,
			q,
			s.reporter,
			not_found.New(constant.DecisionKind, raw),
		)

		return nil, false
	}

	d, f := s.service.Decision(uint(identifier))

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return nil, false
	}

	return d, true
}
