package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestServiceUpdateMemoryPreservesTags(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "pace"
	p.Content = "original content"
	p.Description = "original desc"
	p.Type = "feedback"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		o.Service.AddTags(
			m.Identifier,
			[]string{"always", "go-conventions"},
		),
	)
	q := save_option.New()
	q.Name = "pace"
	q.Content = "updated content"
	q.Description = "updated desc"
	q.Source = "test"
	updated, e := o.Service.UpdateMemory(m.Identifier, q)
	assert.FatalOnError(t, e)
	assert.String(t, "updated content", updated.Content)
	assert.Count(t, 2, updated.Tags)
	assert.Count(t, 3, o.Indexer.Pushed)
}

func TestServiceUpdateMemoryNonexistentFails(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = constant.FixtureName
	p.Content = constant.FixtureContent
	p.Description = "desc"
	p.Source = "test"
	_, e := o.Service.UpdateMemory(999, p)
	assert.Error(t, e)
}
