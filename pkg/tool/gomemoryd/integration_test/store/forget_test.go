//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestForgetMemory(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "to forget"
	o.Description = "desc"
	o.Type = "feedback"
	identifier := s.CreateMemory(o)
	s.ForgetMemory(identifier, "test")
	assert.Count(t, 0, s.ListMemories("", "", true))
	assert.Count(t, 1, s.ListMemories("", "", false))
}
