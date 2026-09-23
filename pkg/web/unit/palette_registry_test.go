package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"testing"
)

func TestRegistrySearchEmpty(t *testing.T) {
	r := web_tester.NewRegistry()
	results := r.Search("")
	assert.Integer(t, 7, len(results))
}

func TestRegistrySearchFilters(t *testing.T) {
	r := web_tester.NewRegistry()
	results := r.Search("sessions")
	assert.Integer(t, 1, len(results))
}

func TestRegistrySearchNoResults(t *testing.T) {
	r := web_tester.NewRegistry()
	results := r.Search("xyz")
	assert.Integer(t, 0, len(results))
}

func TestRegistrySearchRanking(t *testing.T) {
	r := web_tester.NewRegistry()
	results := r.Search("sb")
	assert.True(t, len(results) > 0)
	assert.String(t, "Start build", results[0].Command.Label)
}

func TestRegistrySearchAcronym(t *testing.T) {
	r := web_tester.NewRegistry()
	results := r.Search("cp")
	assert.True(t, len(results) > 0)
	assert.String(t, "Create project", results[0].Command.Label)
}

func TestRegistrySearchSingleCharacter(t *testing.T) {
	r := web_tester.NewRegistry()
	results := r.Search("s")
	assert.True(t, len(results) >= 3)
}
