package mattermost_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/integration_test/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"strconv"
	"testing"
	"time"
)

// Pins the chunk-forward behavior: the server caps posts-since
// responses keeping the oldest, so a capped first page must be
// followed up past its newest post and the pages merged.
func TestPostsSinceChunks(t *testing.T) {
	capped := &model.PostList{Posts: map[string]*model.Post{}}

	for i := range constant.MattermostSinceChunkThreshold {
		identifier := fmt.Sprintf("bravo%d", i)
		capped.Order = append(capped.Order, identifier)
		capped.Posts[identifier] = &model.Post{
			Id:       identifier,
			UserId:   "delta",
			Message:  "chunk one",
			CreateAt: int64(2000 + i),
		}
	}

	rest := &model.PostList{
		Order: []string{"foxtrot"},
		Posts: map[string]*model.Post{
			"foxtrot": {
				Id:       "foxtrot",
				UserId:   "delta",
				Message:  "chunk two",
				CreateAt: 9000,
			},
		},
	}
	var calls []string
	r := mattermost_client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/channels/alfa/posts",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					since := q.URL.Query().Get("since")
					calls = append(calls, since)
					milli, e := strconv.ParseInt(since, 10, 64)

					if e != nil {
						t.Error(e)
					}

					if milli < 2999 {
						web.Encode(w, capped)

						return
					}

					web.Encode(w, rest)
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
	posts, e := r.Client.PostsSince(
		&model.Channel{Id: "alfa"},
		time.UnixMilli(1000),
	)

	if e != nil {
		t.Fatal(e)
	}

	assert.Any(t, 2, len(calls))
	assert.Any(t, len(capped.Order)+1, len(posts))
	assert.Any(t, "bravo0", posts[0].Raw.Id)
	assert.Any(t, "foxtrot", posts[len(posts)-1].Raw.Id)
}
