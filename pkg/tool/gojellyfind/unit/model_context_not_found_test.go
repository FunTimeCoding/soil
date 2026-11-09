package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"testing"
)

func TestUnknownItemIsAnsweredQuietlyOverModelContext(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := s.ContextServer.Connect(t)
	defer c.Close()
	assert.String(
		t,
		"item not found: charlie",
		c.MustCallToolError(constant.GetItem, map[string]any{"id": "charlie"}),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
