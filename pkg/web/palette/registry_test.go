package palette

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func testRegistry() *Registry {
	r := NewRegistry()
	r.Register(
		Command{Label: "Dashboard", Path: "/", Category: "navigate"},
		Command{
			Label:    "Create project",
			Path:     "/projects/new",
			Category: "action",
		},
		Command{Label: "Sessions", Path: "/sessions", Category: "navigate"},
		Command{
			Label:    "Start build",
			Path:     "/builds/start",
			Category: "action",
		},
		Command{Label: "Metrics", Path: "/metrics", Category: "navigate"},
		Command{
			Label:    "Push deploy",
			Path:     "/deploys/push",
			Category: "action",
		},
		Command{
			Label:    "Search logs",
			Path:     "/logs/search",
			Category: "navigate",
		},
	)

	return r
}

func TestRegistrySearchEmpty(t *testing.T) {
	r := testRegistry()
	results := r.Search("")
	assert.Integer(t, 7, len(results))
}

func TestRegistrySearchFilters(t *testing.T) {
	r := testRegistry()
	results := r.Search("sessions")
	assert.Integer(t, 1, len(results))
}

func TestRegistrySearchNoResults(t *testing.T) {
	r := testRegistry()
	results := r.Search("xyz")
	assert.Integer(t, 0, len(results))
}

func TestRegistrySearchRanking(t *testing.T) {
	r := testRegistry()
	results := r.Search("sb")
	assert.True(t, len(results) > 0)
	assert.String(t, "Start build", results[0].Command.Label)
}

func TestRegistrySearchAcronym(t *testing.T) {
	r := testRegistry()
	results := r.Search("cp")
	assert.True(t, len(results) > 0)
	assert.String(t, "Create project", results[0].Command.Label)
}

func TestRegistrySearchSingleCharacter(t *testing.T) {
	r := testRegistry()
	results := r.Search("s")
	assert.True(t, len(results) >= 3)
}
