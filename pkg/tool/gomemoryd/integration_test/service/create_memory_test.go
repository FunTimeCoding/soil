package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestServiceCreateMemory(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "retry policy"
	p.Content = "Retry failed requests with exponential backoff."
	p.Description = "retry with backoff"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.String(t, "retry policy", m.Name)
	assert.String(t, "feedback", m.Type)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.String(t, "memory/1", o.Indexer.Pushed[0].Name)
}

func TestServiceCreateMemoryWithExplicitType(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "deploy target"
	p.Content = "Production deployments use blue-green strategy."
	p.Description = "blue-green deploy pattern"
	p.Type = "user"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.String(t, "user", m.Type)
}
