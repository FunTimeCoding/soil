//go:build local

package coordination

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"github.com/google/uuid"
	"testing"
)

func TestCheckCreatesSession(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	assert.True(t, a.Name() != "")
}

func TestCheckPreviewNoSession(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	restClient := s.RESTClient(t)
	response, e := restClient.GetCheckWithResponse(
		context.Background(),
		&client.GetCheckParams{
			Session: uuid.New().String(),
			Preview: new(true),
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, response.JSON200.Changed)
	assert.Count(t, 0, response.JSON200.Entries)
}

func TestCheckRecentActivity(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "build search index")
	b.CheckLive()
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "indexing done",
		},
	)
	check := b.CheckLive()
	completions := clientEntriesByKind(
		check.Entries,
		constant.QueueSessionComplete,
	)
	assert.True(t, len(completions) > 0)
	found := false

	for _, c := range completions {
		if c.Body != "" {
			found = true
		}
	}

	assert.True(t, found)
}

func TestCheckRecentActivityLive(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "search index")
	b.Check()
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "indexing done",
		},
	)
	check := b.CheckLive()
	assert.True(t, check.Changed)
	completions := clientEntriesByKind(
		check.Entries,
		constant.QueueSessionComplete,
	)
	assert.True(t, len(completions) > 0)
}

func TestCheckRecentActivityUpdate(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "initial")
	b.CheckLive()
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone reached",
		},
	)
	check := b.CheckLive()
	updates := clientEntriesByKind(check.Entries, constant.QueueSessionUpdate)
	assert.True(t, len(updates) > 0)
}
