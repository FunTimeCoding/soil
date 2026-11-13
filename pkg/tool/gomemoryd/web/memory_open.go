package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/memory_page"
	"net/http"
)

func (s *Server) memoryOpen(
	w http.ResponseWriter,
	q *http.Request,
) {
	v, okay := s.visibleMemory(w, q)

	if !okay {
		return
	}

	s.view.RenderFragment(w, memory_page.New(v).Opened())
}
