package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration_test/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
	"time"
)

func TestChannelPosts(t *testing.T) {
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/teams/tango/channels/name/alfa",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					web.Encode(
						w,
						&model.Channel{
							Id:   "alfa",
							Name: "alfa",
						},
					)
				},
			)
			m.HandleFunc(
				"/api/v4/channels/alfa/posts",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
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
					web.Encode(
						w,
						&model.User{
							Id:       "delta",
							Username: "echo",
						},
					)
				},
			)
		},
	)
	result, e := r.Client.GetChannelPostsWithResponse(
		context.Background(),
		"alfa",
		&client.GetChannelPostsParams{Since: time.UnixMilli(1000)},
	)

	if e != nil {
		t.Fatal(e)
	}

	assert.Any(t, 200, result.StatusCode())
	posts := *result.JSON200
	assert.Any(t, 3, len(posts))
	assert.Any(t, "bravo", posts[0].Identifier)
	assert.Any(t, "echo", posts[0].Username)
	assert.Any(t, "charlie", posts[1].Identifier)
	assert.Any(t, "bravo", *posts[1].Root)
	assert.Any(t, "foxtrot", posts[2].Identifier)
	assert.Any(t, []string{"golf"}, *posts[2].Files)
}
