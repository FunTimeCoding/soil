package mechanic

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
)

func (s *Server) fire(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(constant.TriggerHeader, constant.FiredEvent)
	s.view.RenderFragment(w, gomponents.Text(""))
}
