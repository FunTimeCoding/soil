//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

func TestAddAndListContexts(t *testing.T) {
	s := model_context_tester.New(t)
	s.AddContext("docs", "/design/", "Architecture and design decisions")
	result := s.ListContexts()
	assert.StringContains(t, "design", result)
	assert.StringContains(t, "Architecture", result)
}

func TestRemoveContext(t *testing.T) {
	s := model_context_tester.New(t)
	s.AddContext("docs", "/temp/", "Temporary context")
	s.RemoveContext("docs", "/temp/")
	result := s.ListContexts()
	assert.StringNotContains(t, "Temporary", result)
}
