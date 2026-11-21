package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/upgrade_page"
	"net/http"
)

func (s *Server) upgradeOpen(
	w http.ResponseWriter,
	r *http.Request,
) {
	g, e := s.system.Upgrade(
		r.URL.Query().Get(constant.ReleaseParameter),
		r.URL.Query().Get(constant.SourceParameter),
	)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderFragment(w, upgrade_page.Detail(g))
}
