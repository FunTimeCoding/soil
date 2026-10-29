package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
)

func TestThread(t *testing.T) {
	list := &model.PostList{
		Order: []string{"foxtrot", "charlie", "bravo"},
		Posts: map[string]*model.Post{
			"bravo": {
				Id:         "bravo",
				UserId:     "delta",
				Message:    "root",
				CreateAt:   2000,
				ReplyCount: 2,
			},
			"charlie": {
				Id:       "charlie",
				UserId:   "delta",
				RootId:   "bravo",
				Message:  "reply",
				CreateAt: 3000,
			},
			"foxtrot": {
				Id:       "foxtrot",
				UserId:   "delta",
				RootId:   "bravo",
				CreateAt: 4000,
				FileIds:  []string{"golf"},
			},
		},
	}
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/posts/bravo",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					web.Encode(w, list.Posts["bravo"])
				},
			)
			m.HandleFunc(
				"/api/v4/posts/bravo/thread",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					web.Encode(w, list)
				},
			)
			m.HandleFunc(
				"/api/v4/users/delta",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					web.Encode(w, &model.User{Id: "delta", Username: "echo"})
				},
			)
		},
	)
	replies, e := r.Client.GetThread(context.Background(), "bravo")
	assert.FatalOnError(t, e)
	assert.Any(t, 2, len(replies))
	assert.Any(t, "charlie", replies[0].Identifier)
	assert.Any(t, "foxtrot", replies[1].Identifier)
	assert.Any(t, []string{"golf"}, *replies[1].Files)
}
