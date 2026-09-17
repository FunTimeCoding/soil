package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestMatchRoute(t *testing.T) {
	paths := []string{"/api/alerts", "/api/alerts/{name}/history"}
	assert.True(t, pointer.MatchRoute("/api/alerts", paths))
	assert.True(t, pointer.MatchRoute("/api/alerts/", paths))
	assert.True(t, pointer.MatchRoute("/api/alerts/{id}/history", paths))
	assert.True(t, pointer.MatchRoute("/api/alerts/high/history", paths))
	assert.True(t, pointer.MatchRoute("/api/...", paths))
	assert.False(t, pointer.MatchRoute("/api/alerts/high", paths))
	assert.False(t, pointer.MatchRoute("/api/ghosts", paths))
	assert.False(t, pointer.MatchRoute("/api", paths))
}
