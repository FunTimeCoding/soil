package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) redirect(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(constant.ExtendedRedirect, constant.StreamPath)
	s.view.RenderFragment(w, gomponents.Text(""))
}
