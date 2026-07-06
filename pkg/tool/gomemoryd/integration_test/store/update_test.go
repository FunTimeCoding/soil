//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestUpdateCreatesVersion(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "original"
	o.Description = "desc"
	o.Type = "feedback"
	identifier := s.CreateMemory(o)
	p := save_option.New()
	p.Content = "updated content"
	p.Description = "new desc"
	p.Tags = []string{"always"}
	s.UpdateMemory(identifier, p)
	history := s.GetMemoryHistory(identifier)
	assert.Count(t, 2, history)
	assert.String(t, "created", history[0].ChangeType)
	assert.String(t, "updated", history[1].ChangeType)
	m := s.GetMemory(identifier)
	assert.String(t, "updated content", m.Content)
	assert.Count(t, 1, m.Tags)
	assert.String(t, "always", m.Tags[0])
}
