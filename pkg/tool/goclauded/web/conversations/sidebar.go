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

	sessions, e := s.service.EnrichedSessions(0, 0)
	errors.PanicOnError(e)
	limit := 30

	if skip >= len(sessions) {
		return
	}

	sessions = sessions[skip:]
	hasMore := len(sessions) > limit

	if hasMore {
		sessions = sessions[:limit]
	}

	var entries []gomponents.Node

	for _, e := range sessions {
		entries = append(entries, entry(e))
	}

	if hasMore {
		entries = append(entries, sentinel(skip+limit))
	}

	w.Header().Set(web.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(gomponents.Group(entries).Render(w))
}
