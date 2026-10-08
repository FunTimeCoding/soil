package coordination

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"path/filepath"
	"testing"
)

func TestRestSearchConversations(t *testing.T) {
	s := base.New(t)
	s.WriteSessionFile("alfa", "alfa-slug")
	s.Service.PopulateCache()
	s.Service.BackfillSessions()
	s.Service.CheckConsistency()
	s.Service.CatchUpSearch()
	r, e := s.RESTClient(t).GetSessionsSearchWithResponse(
		context.Background(),
		&client.GetSessionsSearchParams{Query: "login bug"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.Integer(t, 1, r.JSON200.Indexed)
	assert.Integer(t, 1, r.JSON200.Total)
	assert.Integer(t, 1, len(r.JSON200.Conversations))
	c := r.JSON200.Conversations[0]
	assert.String(t, "alfa", c.Session)
	assert.String(t, "alfa-slug", c.Name)
	assert.Integer(t, 1, c.Count)
	assert.String(t, "message", c.Hits[0].Kind)
	assert.String(t, "can you help me fix the login bug", c.Hits[0].Snippet)
}

func TestRestSessionWindow(t *testing.T) {
	s := base.New(t)
	system.WriteFile(
		filepath.Join(s.Harbor, "charlie.jsonl"),
		[]byte(
			`{"uuid":"c1","type":"user","timestamp":"2026-10-08T01:00:00Z","message":{"role":"user","content":"first question"}}
{"uuid":"c2","type":"assistant","timestamp":"2026-10-08T01:01:00Z","message":{"role":"assistant","content":[{"type":"text","text":"first answer"}]}}
`,
		),
		0644,
	)
	s.Service.CatchUpSearch()
	c := s.RESTClient(t)
	r, e := c.GetSessionWindowWithResponse(
		context.Background(),
		"charlie",
		&client.GetSessionWindowParams{Around: "c2"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.Integer(t, 2, len(r.JSON200.Blocks))
	assert.String(t, "user", r.JSON200.Blocks[0].Role)
	assert.String(t, "first answer", r.JSON200.Blocks[1].Text)
	missing, e := c.GetSessionWindowWithResponse(
		context.Background(),
		"charlie",
		&client.GetSessionWindowParams{Around: "nope"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 404, missing.StatusCode())
	assert.String(
		t,
		"no block nope in conversation charlie",
		missing.JSON404.Error,
	)
}

func TestRestSearchEmptyIsArray(t *testing.T) {
	s := base.New(t)
	r, e := s.RESTClient(t).GetSessionsSearchWithResponse(
		context.Background(),
		&client.GetSessionsSearchParams{Query: "nothing indexed"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.StringContains(t, `"conversations":[]`, string(r.Body))
}
