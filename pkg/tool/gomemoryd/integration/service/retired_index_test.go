package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestRetiredMemoryStaysOutOfTheIndexOnTagEdit(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "alfa"
	p.Content = "first"
	p.Description = "a test"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	pushed := len(o.Indexer.Pushed)
	assert.FatalOnError(
		t,
		o.Service.ReplaceTags(m.Identifier, []string{"build"}),
	)
	assert.Integer(t, pushed, len(o.Indexer.Pushed))
	assert.String(
		t,
		"memory/1",
		o.Indexer.Deleted[len(o.Indexer.Deleted)-1].Path,
	)
}

func TestActiveMemoryIsIndexedOnTagEdit(t *testing.T) {
	o := service_tester.New(t)
	p := save_option.New()
	p.Name = "bravo"
	p.Content = "second"
	p.Description = "a test"
	p.Source = "test"
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)
	pushed := len(o.Indexer.Pushed)
	assert.FatalOnError(
		t,
		o.Service.ReplaceTags(m.Identifier, []string{"build"}),
	)
	assert.Integer(t, pushed+1, len(o.Indexer.Pushed))
}
