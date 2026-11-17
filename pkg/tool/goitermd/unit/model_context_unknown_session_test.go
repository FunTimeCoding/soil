package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/base"
	"testing"
)

func TestKeyToUnknownSessionIsAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.String(
		t,
		"step 1: session alfa not found",
		c.MustCallToolError(
			constant.SendKey,
			map[string]any{"session_id": "alfa", "keys": []string{"Return"}},
		),
	)
}
