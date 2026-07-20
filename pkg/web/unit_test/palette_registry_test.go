package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"testing"
)

func testRegistry() *palette.Registry {
	r := palette.NewRegistry()
	r.Register(
		palette.Command{Label: "Dashboard", Path: "/", Category: "navigate"},
		palette.Command{
			Label:    "Create project",
			Path:     "/projects/new",
			Category: "action",
		},
		palette.Command{
			Label:    "Sessions",
			Path:     "/sessions",
			Category: "navigate",
		},
		palette.Command{
			Label:    "Start build",
			Path:     "/builds/start",
			Category: "action",
		},
		palette.Command{
			Label:    "Metrics",
			Path:     "/metrics",
			Category: "navigate",
		},
		palette.Command{
			Label:    "Push deploy",
			Path:     "/deploys/push",
			Category: "action",
		},
		palette.Command{
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
