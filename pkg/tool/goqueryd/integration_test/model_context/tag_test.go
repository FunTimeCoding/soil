//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

func TestTagSetAndGet(t *testing.T) {
	s := model_context_tester.New(t)
	s.SetTag("design/", "design-doc")
	result := s.GetTag("design/")
	assert.StringContains(t, "design-doc", result)
}

func TestTagRemove(t *testing.T) {
	s := model_context_tester.New(t)
	s.SetTag("temp/", "scratch")
	s.SetTag("temp/", "")
	result := s.GetTag("temp/")
	assert.StringContains(t, "no tag", result)
}
