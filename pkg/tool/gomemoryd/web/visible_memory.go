package web

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_view"
	"net/http"
	"strconv"
)

func (s *Server) visibleMemory(
	w http.ResponseWriter,
	q *http.Request,
) (*memory_view.View, bool) {
	raw := q.PathValue(constant.Identifier)
	identifier, e := strconv.ParseInt(raw, 10, 64)

	if e != nil {
		s.view.Answer(w, q, s.reporter, not_found.New(constant.MemoryKind, raw))

		return nil, false
	}

	v, f := s.service.MemoryView(identifier)

	if f != nil {
		s.view.Answer(w, q, s.reporter, f)

		return nil, false
	}

	return v, true
}
