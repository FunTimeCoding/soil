//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration/model_context_tester"
	"testing"
)

func indexFixtures(t *testing.T) *model_context_tester.Tester {
	t.Helper()
	s := model_context_tester.New(t)
	s.IndexFixtures()

	return s
}
