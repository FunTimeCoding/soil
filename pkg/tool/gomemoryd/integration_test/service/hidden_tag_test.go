package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestHiddenTagLeavesIndex(t *testing.T) {
	o := service_tester.New(t)
	o.Service.WithHiddenTag("private")
	p := save_option.New()
	p.Name = "delta"
	p.Content = "quiet content"
	p.Description = "quiet description"
	p.Type = "project"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.FatalOnError(t, o.Service.AddTags(m.Identifier, []string{"private"}))
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.Count(t, 1, o.Indexer.Deleted)
	q := save_option.New()
	q.Name = "delta"
	q.Content = "updated quiet content"
	q.Description = "quiet description"
	q.Source = "test"
	_, e = o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.Count(t, 2, o.Indexer.Deleted)
	assert.FatalOnError(
		t,
		o.Service.RemoveTags(m.Identifier, []string{"private"}),
	)
	assert.Count(t, 2, o.Indexer.Pushed)
	hidden, e := o.Service.HiddenIdentifiers()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, hidden)
}
