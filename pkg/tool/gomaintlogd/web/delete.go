package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web/entry_page"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := queryIdentifier(r)

	if e == nil {
		e = s.service.Delete(identifier)
	}

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	s.view.RenderStatusFragment(
		w,
		entry_page.Removal(identifier),
		"Deleted entry",
		webConstant.LayoutStatusSuccess,
	)
}
