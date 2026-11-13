package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/permission_page"
	"net/http"
)

func (s *Server) permissionOpen(
	w http.ResponseWriter,
	q *http.Request,
) {
	r, e := s.transcript.Permissions(constant.PermissionWindowDays)

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	blocker := q.URL.Query().Get(constant.PermissionBlocker)
	row := permissionRow(r, blocker)

	if row == nil {
		s.view.RenderFragment(w, permission_page.Removal(blocker))

		return
	}

	s.view.RenderFragment(w, permission_page.Detail(row))
}
