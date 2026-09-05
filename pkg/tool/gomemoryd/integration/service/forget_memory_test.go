package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestServiceForgetMemory(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "to forget"
	p.Content = constant.FixtureContent
	p.Description = "desc"
	p.Type = "feedback"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	assert.Count(t, 1, o.Indexer.Deleted)
	assert.String(t, "memory/1", o.Indexer.Deleted[0].Path)
	assert.String(t, "memories", o.Indexer.Deleted[0].Collection)
	active, e := o.Service.ListMemories("", "", "", true)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, active)
}
