package request_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"strconv"
)

func (s *Server) serve(
	w http.ResponseWriter,
	q *http.Request,
) {
	s.lock.Lock()
	s.queries = append(s.queries, q.URL.Query())
	s.lock.Unlock()
	page, e := strconv.Atoi(q.URL.Query().Get("page"))
	errors.PanicOnError(e)
	size, e := strconv.Atoi(q.URL.Query().Get("per_page"))
	errors.PanicOnError(e)
	list := []map[string]any{}

	for i := (page-1)*size + 1; i <= min(page*size, s.total); i++ {
		list = append(
			list,
			map[string]any{
				"iid":        i,
				"project_id": 7,
				"title":      join.Empty("change ", strconv.Itoa(i)),
				"state":      "opened",
				"created_at": "2026-11-20T10:00:00Z",
				"updated_at": "2026-11-25T10:00:00Z",
				"author":     map[string]any{"username": "alfa"},
				"reviewers":  []map[string]any{{"username": "bravo"}},
				"assignees": []map[string]any{
					{"username": "charlie"},
				},
				"source_branch":         "feature",
				"target_branch":         "main",
				"draft":                 true,
				"detailed_merge_status": "ci_still_running",
				"user_notes_count":      3,
			},
		)
	}

	web.Encode(w, list)
}
