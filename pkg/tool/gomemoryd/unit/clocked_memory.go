package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func clockedMemory(
	t *testing.T,
	s *base.Server,
) int64 {
	t.Helper()
	m, e := s.MemoryClient(t).PostMemory(
		t.Context(),
		client.PostMemoryJSONRequestBody{
			Name:        "clocked",
			Content:     "clocked content",
			Description: "clocked memory",
			Type:        new("feedback"),
		},
	)
	assert.FatalOnError(t, e)

	return m.Identifier
}
