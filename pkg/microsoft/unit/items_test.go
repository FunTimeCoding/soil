package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_tester"
	"github.com/funtimecoding/soil/pkg/microsoft/mock_token_store"
	"testing"
)

func TestItemsFollowNextLink(t *testing.T) {
	g := graph_tester.New(t)
	s := g.Paged()
	items, e := g.Client(
		s,
		mock_token_store.New(t, graph_tester.Valid()),
	).Items("/v1.0/items", nil, 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 4, items)
	next := s.Exchanges("/v1.0/next")
	assert.Count(t, 1, next)
	assert.String(t, "2", next[0].Query.Get("skip"))
	assert.String(t, "Bearer stored-access", next[0].Authorization)
}

func TestItemsStopAtLimit(t *testing.T) {
	g := graph_tester.New(t)
	s := g.Paged()
	items, e := g.Client(
		s,
		mock_token_store.New(t, graph_tester.Valid()),
	).Items("/v1.0/items", nil, 2)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, items)
	assert.Count(t, 0, s.Exchanges("/v1.0/next"))
}

func TestItemsTruncateToLimit(t *testing.T) {
	g := graph_tester.New(t)
	items, e := g.Client(
		g.Paged(),
		mock_token_store.New(t, graph_tester.Valid()),
	).Items("/v1.0/items", nil, 3)
	assert.FatalOnError(t, e)
	assert.Count(t, 3, items)
	assert.String(t, "3", string(items[2]))
}
