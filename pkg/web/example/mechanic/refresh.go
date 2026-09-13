package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) refresh(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(constant.RefreshHeader, "true")
	s.view.RenderFragment(w, gomponents.Text(""))
}
