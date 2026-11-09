package unit

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"slices"
	"sync"
)

func threadChannel(
	deleted *[]string,
	guard *sync.Mutex,
	posts ...*model.Post,
) func(*http.ServeMux) {
	live := slices.Clone(posts)

	return func(m *http.ServeMux) {
		m.HandleFunc(
			"GET /api/v4/users/me",
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				web.Encode(w, &model.User{Id: "bot"})
			},
		)
		m.HandleFunc(
			"GET /api/v4/users/{identifier}",
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				web.Encode(
					w,
					&model.User{
						Id:       q.PathValue("identifier"),
						Username: q.PathValue("identifier"),
					},
				)
			},
		)
		m.HandleFunc(
			"GET /api/v4/channels/echo/posts",
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				guard.Lock()
				defer guard.Unlock()
				l := model.NewPostList()

				if q.URL.Query().Get("before") != "" {
					web.Encode(w, l)

					return
				}

				collapsed := q.URL.Query().Get("collapsedThreads") == "true"

				for i := len(live) - 1; i >= 0; i-- {
					p := live[i]

					if collapsed && p.RootId != "" {
						continue
					}

					l.AddPost(p)
					l.AddOrder(p.Id)
				}

				web.Encode(w, l)
			},
		)
		m.HandleFunc(
			"DELETE /api/v4/posts/{identifier}",
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				guard.Lock()
				defer guard.Unlock()
				identifier := q.PathValue("identifier")
				*deleted = append(*deleted, identifier)
				live = slices.DeleteFunc(
					live,
					func(p *model.Post) bool { return p.Id == identifier },
				)
				web.Encode(w, map[string]string{"status": "OK"})
			},
		)
	}
}
