//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

func TestAddCollectionAndIndex(t *testing.T) {
	s := model_context_tester.New(t)
	s.IndexFixtures()
	result := s.MustCallTool(constant.Index, map[string]any{})
	assert.StringContains(t, "test", result)
}

func TestStatusAfterIndex(t *testing.T) {
	s := model_context_tester.New(t)
	s.IndexFixtures()
	result := s.MustCallTool(constant.Status, map[string]any{})
	assert.StringContains(t, "total_documents", result)
}
