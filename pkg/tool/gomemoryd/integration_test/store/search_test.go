//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestSearchMemories(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "Always validate input at service boundaries"
	o.Description = "validate at boundaries; trust internal callers"
	o.Type = "feedback"
	o.Tags = []string{"always"}
	s.CreateMemory(o)
	p := save_option.New()
	p.Content = "Build regression tests to verify fix correctness"
	p.Description = "regression tests over manual verification"
	p.Type = "feedback"
	s.CreateMemory(p)
	results := s.SearchMemories("validate input", 10, "", "")
	assert.Greater(t, 0, float64(len(results)))
	assert.String(
		t,
		"validate at boundaries; trust internal callers",
		results[0].Description,
	)
}
