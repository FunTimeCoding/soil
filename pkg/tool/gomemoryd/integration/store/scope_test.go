package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestScopeSeparatesListing(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "default memory"
	o.Content = "Lives in the default scope."
	o.Description = "default"
	o.Type = "feedback"
	s.CreateMemory(o)
	p := save_option.New()
	p.Name = "scoped memory"
	p.Content = "Lives in a named scope."
	p.Description = "scoped"
	p.Type = "reference"
	p.Scope = "alpha"
	s.CreateMemory(p)
	defaults := s.ListMemories("", "", "", true)
	assert.Count(t, 1, defaults)
	assert.String(t, "default memory", defaults[0].Name)
	scoped := s.ListMemories("", "", "alpha", true)
	assert.Count(t, 1, scoped)
	assert.String(t, "scoped memory", scoped[0].Name)
	assert.String(t, "alpha", scoped[0].Scope)
	all := s.ListMemories("", "", constant.AllScope, true)
	assert.Count(t, 2, all)
}

func TestScopeSeparatesSearch(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Name = "default note"
	o.Content = "The turbine spins in the default scope."
	o.Description = "default turbine"
	o.Type = "feedback"
	s.CreateMemory(o)
	p := save_option.New()
	p.Name = "scoped note"
	p.Content = "The turbine spins in a named scope."
	p.Description = "scoped turbine"
	p.Type = "reference"
	p.Scope = "alpha"
	s.CreateMemory(p)
	defaults := s.SearchMemories("turbine", 10, "", "", "")
	assert.Count(t, 1, defaults)
	assert.String(t, "default note", defaults[0].Name)
	scoped := s.SearchMemories("turbine", 10, "", "", "alpha")
	assert.Count(t, 1, scoped)
	assert.String(t, "scoped note", scoped[0].Name)
	all := s.SearchMemories("turbine", 10, "", "", constant.AllScope)
	assert.Count(t, 2, all)
}
