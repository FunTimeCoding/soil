package conversations

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
	"strconv"
)

func (s *Server) sidebar(
	w http.ResponseWriter,
	r *http.Request,
) {
	skip := 0

	if v := r.URL.Query().Get(constant.Offset); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			skip = n
		}
	}

	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(gomponents.Group(s.sidebarNodes(skip)).Render(w))
}
