package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestListRelated(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "error handling"
	o.Content = "error handling content"
	o.Description = "error handling description"
	o.Type = "feedback"
	id1 := s.CreateMemory(o)
	p := save_option.New()
	p.Name = "deployment"
	p.Content = "deployment content"
	p.Description = "deployment description"
	p.Type = "feedback"
	p.Tags = []string{"deploy", "no-index"}
	id2 := s.CreateMemory(p)
	q := save_option.New()
	q.Name = "telemetry"
	q.Content = "telemetry content"
	q.Description = "telemetry description"
	q.Type = "project"
	id3 := s.CreateMemory(q)
	s.CreateRelation(id1, id2)
	s.CreateRelation(id3, id1)
	related := s.ListRelated(id1)
	assert.Count(t, 2, related)
	assert.Integer64(t, id2, related[0].Identifier)
	assert.String(t, "deployment", related[0].Name)
	assert.String(t, "deployment description", related[0].Description)
	assert.Count(t, 2, related[0].Tags)
	assert.Integer64(t, id3, related[1].Identifier)
	assert.String(t, "telemetry", related[1].Name)
	s.ForgetMemory(id2, "test")
	assert.Count(t, 1, s.ListRelated(id1))
}
