package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
)

func TestEditPostPatchesOnlyTheMessage(t *testing.T) {
	var sent map[string]any
	c := mattermost_client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"PUT /api/v4/posts/alfa/patch",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					if e := json.NewDecoder(q.Body).Decode(&sent); e != nil {
						t.Error(e)
					}

					web.Encode(
						w,
						&model.Post{Id: "alfa", Message: "done: bravo"},
					)
				},
			)
		},
	).Client
	p, e := c.EditPost(&model.Post{Id: "alfa"}, "done: bravo")
	assert.FatalOnError(t, e)
	assert.String(t, "done: bravo", p.Message)
	assert.Any(t, "done: bravo", sent["message"])
	assert.Any(t, nil, sent["props"])
	assert.Any(t, nil, sent["is_pinned"])
}

func TestEditingAMissingPostIsNotFound(t *testing.T) {
	_, e := mattermost_client_tester.New(t, statusUpstream).Client.EditPost(
		&model.Post{Id: "missing"},
		"done",
	)
	assert.True(t, not_found.Is(e))
	assert.String(t, "post not found: missing", e.Error())
}

func TestTheMockEditsItsStoredPost(t *testing.T) {
	m := mock_client.New("charlie")
	defer m.Close()
	m.AddPost("alfa")
	_, e := m.EditPost(&model.Post{Id: "alfa"}, "done: bravo")
	assert.FatalOnError(t, e)
	p, f := m.FindPost("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "done: bravo", p.Message)
	_, g := m.EditPost(&model.Post{Id: "delta"}, "done")
	assert.True(t, not_found.Is(g))
}
