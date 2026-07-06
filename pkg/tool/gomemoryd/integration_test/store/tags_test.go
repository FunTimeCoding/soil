//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestListTags(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "a"
	o.Description = "desc a"
	o.Type = "feedback"
	o.Tags = []string{"always", "deployment"}
	s.CreateMemory(o)
	p := save_option.New()
	p.Content = "b"
	p.Description = "desc b"
	p.Type = "feedback"
	p.Tags = []string{"always"}
	s.CreateMemory(p)
	tags := s.ListTags()
	assert.Count(t, 2, tags)
	assert.String(t, "always", tags[0].Tag)
	assert.Integer(t, 2, tags[0].Count)
}
