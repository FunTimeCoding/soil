package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/permission_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) permissionsPage(
	w http.ResponseWriter,
	q *http.Request,
) {
	r, e := s.transcript.Permissions(constant.PermissionWindowDays, time.Now())

	if e != nil {
		s.view.Answer(w, q, s.reporter, e)

		return
	}

	s.view.RenderPageWithSummary(
		w,
		constant.PermissionsTitle,
		constant.PermissionsPath,
		permission_page.Summary(r),
		html.H3(gomponents.Text(constant.PermissionsTitle)),
		permission_page.Region(r),
	)
}
