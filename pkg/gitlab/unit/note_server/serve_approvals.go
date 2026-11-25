package note_server

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) serveApprovals(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.Encode(
		w,
		map[string]any{
			"iid":        3,
			"project_id": 7,
			"approved_by": []map[string]any{
				{"user": map[string]any{"username": "echo"}},
				{"user": map[string]any{"username": "delta"}},
			},
		},
	)
}
