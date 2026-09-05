package mattermost_client

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
	"time"
)

// Pins two regressions: GetPostsSince running with collapsed
// threads (only roots returned), and attachment-only posts with
// an empty message being dropped.
func TestPostsSinceIncludesRepliesAndAttachments(t *testing.T) {
	var collapsed string
	r := mattermost_client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/channels/alfa/posts",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					collapsed = q.URL.Query().Get("collapsedThreads")
					web.Encode(
						w,
						&model.PostList{
							Order: []string{"foxtrot", "charlie", "bravo"},
							Posts: map[string]*model.Post{
								"bravo": {
									Id:       "bravo",
									UserId:   "delta",
									Message:  "root",
									CreateAt: 2000,
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
						},
					)
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

	assert.Any(t, "false", collapsed)
	assert.Any(t, 3, len(posts))
	assert.Any(t, "bravo", posts[0].Raw.Id)
	assert.Any(t, "charlie", posts[1].Raw.Id)
	assert.Any(t, "bravo", posts[1].Raw.RootId)
	assert.Any(t, "foxtrot", posts[2].Raw.Id)
	assert.Any(t, []string{"golf"}, []string(posts[2].Raw.FileIds))
}
