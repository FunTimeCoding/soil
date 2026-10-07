package mattermost_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestPostsSinceFollowsPastACappedPageAndMerges(t *testing.T) {
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

func TestPostsSinceIncludesRepliesAndAttachmentOnlyPosts(t *testing.T) {
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

func TestRefreshSocketReturnsTheDialError(t *testing.T) {
	r := mattermost_client_tester.New(t, func(_ *http.ServeMux) {})
	before := r.Client.WebSocket()
	defer before.Close()
	r.Refuse(1)
	e := r.Client.RefreshSocket()
	assert.NotNil(t, e)
	assert.StringContains(t, "bad handshake", e.Error())
	assert.True(t, before == r.Client.WebSocket())
	assert.Integer(t, 1, r.Accepted())
	assert.Nil(t, r.Client.RefreshSocket())
	after := r.Client.WebSocket()
	defer after.Close()
	assert.True(t, before != after)
	assert.Integer(t, 2, r.Accepted())
}

func TestSocketDeliversPushedEvent(t *testing.T) {
	r := mattermost_client_tester.New(t, func(_ *http.ServeMux) {})
	w := r.Client.WebSocket()
	w.Listen()
	defer w.Close()
	v := &model.WebSocketEvent{}
	v = v.SetEvent(model.WebsocketEventPosted)
	r.Push(
		v.SetData(
			map[string]any{
				constant.MattermostPostField: notation.Encode(
					&model.Post{
						Id:      "bravo",
						RootId:  "alfa",
						UserId:  "delta",
						Message: "pushed from the tester",
					},
					false,
				),
			},
		),
	)

	select {
	case received := <-w.EventChannel:
		assert.String(t, "posted", string(received.EventType()))
		decoded := post.Decode(received)
		assert.String(t, "bravo", decoded.Id)
		assert.String(t, "alfa", decoded.RootId)
		assert.String(t, "pushed from the tester", decoded.Message)
	case <-time.After(constant.MattermostSocketWait):
		t.Fatal("no event received")
	}
}
